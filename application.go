// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package ogwc

import (
	"github.com/c-mueller/ogwc/core"
	"github.com/c-mueller/ogwc/repo"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/op/go-logging"
	"regexp"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type calculationCreationResponse struct {
	Code          int    `json:"code"`
	CalculationID string `json:"calculation_id"`
}

type MetricsUserAccount struct {
	Username string
	Password string
}

var log = logging.MustGetLogger("application")

var crRegex = regexp.MustCompile("cr-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")
var rrRegex = regexp.MustCompile("rr-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type OGWCApplication struct {
	repo   repo.Repository
	engine *gin.Engine
	api    core.OGameAPI

}

func (a *OGWCApplication) Init(c *redis.Options) error {
	a.repo = repo.Repository{
		Options: *c,
	}
	a.api = core.OGAPIRestAPI{}

	a.engine = gin.Default()

	a.engine.POST("/api/v1/submit/:key", a.newCalculation)

	a.engine.GET("/api/v1/calculation/:id", a.getCalculation)
	a.engine.GET("/api/v1/calculation/:id/report", a.getReport)

	a.engine.POST("/api/v1/calculation/:id/add/:key", a.addKey)

	return nil
}

func (a *OGWCApplication) Serve(bindUrl string) {
	a.repo.Connect()
	a.engine.Run(bindUrl)
}

