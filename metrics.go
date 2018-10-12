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
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var requestSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
	Name:      "http_requests",
	Namespace: "ogwc",
	Help:      "Request processing times by Type and ResponseCode",
}, []string{"path", "code"})

func init() {
	prometheus.MustRegister(requestSummary)
}

func (a *OGWCApplication) metricsMiddleware(ctx *gin.Context) {
	start := time.Now()

	ctx.Next()

	execTime := time.Now().Sub(start)

	path := []byte(ctx.Request.URL.Path)

	path = reportRegex.ReplaceAll(path, []byte("REPORT-KEY"))
	path = uidRegex.ReplaceAll(path, []byte("CALCULATION-UID"))

	status := ctx.Writer.Status()

	requestSummary.WithLabelValues(string(path), fmt.Sprintf("%d", status)).Observe(float64(execTime.Nanoseconds()) / 1000000)
}
