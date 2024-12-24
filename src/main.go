package main

import (
	"net/http"
	"time"

	"github.com/InsaneProgZ/user-service/src/adapters/input/controller"
	"github.com/InsaneProgZ/user-service/src/adapters/output/repository"
	"github.com/InsaneProgZ/user-service/src/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func configPrometheus() gin.HandlerFunc {
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)

	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)

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
	routerConfig()
}

func routerConfig() {
	router := gin.Default()

	router.Use(configPrometheus())

	userController := appConfig()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := router.Group("/v1")
	v1.POST("/users", userController.CreateUsers)
	v1.GET("/users/:username", userController.FindUser)

	router.Use(gin.Logger())

	router.Run("localhost:8080")
}

func appConfig() *controller.UserController {
	userRepository := repository.NewUserRepository()
	userPort := service.NewUserService(userRepository)
	userController := controller.NewUserController(userPort)
	return userController
}
