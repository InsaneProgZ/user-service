package main

import (
	"net/http"
	"time"

	"github.com/InsaneProgZ/user-service/src/adapters/controller"
	"github.com/InsaneProgZ/user-service/src/adapters/repository"
	"github.com/InsaneProgZ/user-service/src/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

// PrometheusMiddleware records metrics for Prometheus
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start).Seconds()

		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), http.StatusText(status)).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, c.FullPath(), http.StatusText(status)).Observe(duration)
	}
}

func main() {
	defer glog.Flush()
	router := gin.Default()

	router.Use(PrometheusMiddleware())

	userController := appConfig()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// Defining API routes
	v1 := router.Group("/v1")
	v1.POST("/users", userController.CreateUsers)
	v1.GET("/users/:username", userController.FindUser)

	router.Use(gin.Logger())
	// Run the server
	router.Run(":8080")
}

func appConfig() *controller.UserController {
	defer glog.Flush()
	userRepository := repository.NewUserRepository()
	userPort := service.NewUserService(userRepository)
	userController := controller.NewUserController(userPort)
	return userController
}
