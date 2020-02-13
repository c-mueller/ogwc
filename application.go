// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2019 Christian Müller <dev@c-mueller.xyz>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package ogwc

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/GeertJohan/go.rice"
	"github.com/c-mueller/ogwc/core"
	"github.com/c-mueller/ogwc/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/op/go-logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
)

var log = logging.MustGetLogger("application")

type OGWCApplication struct {
	repo           repo.Repo
	engine         *gin.Engine
	api            core.OGameAPI
	UserAccounts   []MetricsUserAccount
	APIUrlTemplate string
	versionInfo    VersionInfo
}

type VersionInfo struct {
	BuildContext    string `json:"build_context"`
	BuildTimestamp  string `json:"build_timestamp"`
	Revision        string `json:"build_revision"`
	Version         string `json:"version"`
	FrontendHashsum string `json:"frontend_hashsum"`
}

func (a *OGWCApplication) InitVersionInfo(ctx, rev, ver, ts string) {
	a.versionInfo.BuildContext = ctx
	a.versionInfo.Revision = rev
	a.versionInfo.Version = ver
	a.versionInfo.BuildTimestamp = ts

	hashsum := GetUIRevision()
	a.versionInfo.FrontendHashsum = hashsum
}

func GetUIRevision() string {
	ui, err := rice.FindBox("app-ui")
	if err != nil {
		return "DEV-VERSION"
	}
	h := sha256.New()
	ui.Walk("/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			data, err := ui.Bytes(path)
			if err != nil {
				return err
			}
			h.Write(data)
		}

		return nil
	})
	hashsum := make([]byte, 0)
	hashsum = h.Sum(hashsum)
	return hex.EncodeToString(hashsum)
}

func (a *OGWCApplication) Init(useRedis bool, boltpath string, c *redis.Options) error {
	users := a.initMetricsUserAccounts()

	if useRedis {
		a.repo = &repo.RedisRepository{
			Options: *c,
		}
	} else {
		a.repo = &repo.BoltRepository{
			Path: boltpath,
		}
	}
	a.api = core.OGAPIRestAPI{
		QueryURL: a.APIUrlTemplate,
	}

	a.engine = gin.Default()
	a.engine.Use(cors.Default())
	a.engine.Use(a.metricsMiddleware)

	a.initializePrometheusMetricsHandling(users)

	ui, err := rice.FindBox("app-ui")
	if err == nil {
		a.engine.StaticFS("/ui", ui.HTTPBox())

		a.engine.GET("/", a.redirectToUi)
		a.engine.GET("/c/:id", a.redirectToCalculationUi)
		a.engine.GET("/r/:id", a.redirectToCalculationReportUi)
	} else {
		log.Warning("This is a Development Binary. This Means the WebApplication is not available on <URL>/ui")
	}

	a.registerV1ApiMappings()

	return nil
}

func (a *OGWCApplication) redirectToUi(ctx *gin.Context) {
	ctx.Redirect(301, "/ui")
}

func (a *OGWCApplication) redirectToCalculationUi(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.Redirect(301, "/ui/#/calculation/"+id)
}

func (a *OGWCApplication) redirectToCalculationReportUi(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.Redirect(301, "/ui/#/calculation/"+id+"/report")
}

func (a *OGWCApplication) initializePrometheusMetricsHandling(users map[string]string) {
	metricsGroup := a.engine.Group("/metrics", gin.BasicAuth(users))
	prometheusHandler := promhttp.Handler()
	metricsGroup.GET("/", func(ctx *gin.Context) {
		prometheusHandler.ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func (a *OGWCApplication) initMetricsUserAccounts() map[string]string {
	if len(a.UserAccounts) == 0 {
		uid, _ := uuid.NewUUID()
		password := uid.String()
		a.UserAccounts = []MetricsUserAccount{
			{
				Username: "admin",
				Password: password,
			},
		}

		log.Infof("Created Metrics Credentials Username: %q Password: %q", a.UserAccounts[0].Username, a.UserAccounts[0].Password)
	}
	users := make(map[string]string)
	for _, v := range a.UserAccounts {
		users[v.Username] = v.Password
	}

	return users
}

func (a *OGWCApplication) Serve(bindUrl string) {
	err := a.repo.Connect()
	if err != nil {
		log.Error("Connection to Redis failed.")
		log.Error(err.Error())
		os.Exit(1)
	}
	a.engine.Run(bindUrl)
}
