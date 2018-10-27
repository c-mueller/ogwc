// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
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

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type OGWCApplication struct {
	repo         repo.Repository
	engine       *gin.Engine
	api          core.OGameAPI
	UserAccounts []MetricsUserAccount
}

func (a *OGWCApplication) Init(c *redis.Options) error {
	users := a.initMetricsUserAccounts()

	a.repo = repo.Repository{
		Options: *c,
	}
	a.api = core.OGAPIRestAPI{}

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

	a.engine.POST("/api/v1/submit/:key", a.newCalculation)

	a.engine.GET("/api/v1/calculation/:id", a.getCalculation)
	a.engine.GET("/api/v1/calculation/:id/report", a.getReport)
	a.engine.GET("/api/v1/calculation/:id/report/transfers",a.getTransfers)

	a.engine.POST("/api/v1/calculation/:id/add/:key", a.addKey)

	a.engine.POST("/api/v1/calculation/:id/participant/fleet-loss", a.updateAdditionalFleetLoss)
	a.engine.POST("/api/v1/calculation/:id/participant/resource-loss", a.updateAdditionalResourceLoss)
	a.engine.POST("/api/v1/calculation/:id/participant/add", a.addParticipant)
	a.engine.POST("/api/v1/calculation/:id/participant/delete", a.deleteParticipant)

	a.engine.POST("/api/v1/calculation/:id/participant/win/percentage", a.updateWinPercentageOfParticipant)
	a.engine.POST("/api/v1/calculation/:id/participant/win/fixed", a.updateFixedWinOfParticipant)
	a.engine.POST("/api/v1/calculation/:id/participant/win/none", a.updateDisableWinOfParticipant)

	a.engine.POST("/api/v1/calculation/:id/rebalance-win", a.rebalancePercentage)

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
