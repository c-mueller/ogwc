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
