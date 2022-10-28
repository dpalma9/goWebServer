package main

import (
	promMW "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"

	"myweb/metrics"
)

func main() {
	e := echo.New()
	p := promMW.NewPrometheus("echo", nil)

	myCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "my_handler_executions",
		Help:        "Counts executions of my handler function.",
		ConstLabels: prometheus.Labels{"role": "servertest"},
	})
	if err := prometheus.Register(myCounter); err != nil {
		log.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		//myCounter.Inc()
		//return c.JSON(http.StatusOK, "OK")
		return c.JSON(http.StatusOK, metrics.Test(myCounter))
	})

	p.Use(e) // Enable metrics middleware and register route to expose metrics

	if err := e.Start("localhost:8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
