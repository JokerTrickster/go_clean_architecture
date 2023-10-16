package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method"},
	)
	httpRequestsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: []float64{0.1, 0.5, 1, 2, 5, 10},
		},
		[]string{"method"},
	)
)

func PrometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Prometheus 카운터 생성 및 증가
		httpRequestsTotal.WithLabelValues("GET").Inc()
		httpRequestsDuration := prometheus.NewTimer(httpRequestsDuration.WithLabelValues("GET"))
		defer httpRequestsDuration.ObserveDuration()

		return next(c)
	}
}

func init() {
	// Prometheus 지표 등록
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestsDuration)
}
