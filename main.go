package main

import (
	"fmt"

	"./views"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initializeApp()
	router := getEngine()
	router.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}

func getEngine() *gin.Engine {
	router := gin.Default()
	router.Use(gin.ErrorLoggerT(gin.ErrorTypePrivate))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/api/311/v1")
	{
		v1.GET("/services", views.GetServiceList)
		v1.GET("/requests", views.GetServiceRequests)
		v1.POST("/requests", views.ServiceRequest)
		v1.PUT("/images", views.PutRequestImage)
		v1.GET("/images", views.GetRequestImage)
	}

	services := v1.Group("/services")
	{
		services.GET("/:service_code", views.GetServiceDefinition)
	}

	requests := v1.Group("/requests")
	{
		requests.GET("/:service_request_id", views.GetServiceRequestByID)
	}

	return router
}

func initializeApp() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
