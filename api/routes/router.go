package routes

import (
	"context"
	"net/http"

	"github.com/surajkmr91/go-template/commons/middleware"

	"github.com/surajkmr91/go-template/commons/log"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func createRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())
	return router
}

func DefaultRouter(ctx context.Context) *gin.Engine {
	router := createRouter(middleware.Logger())
	//router.GET(c.ActuatorRoute, api.Actuator) // health api

	// v1UserGroup := router.Group(c.V1, middleware.UserAuth())
	// InitLoansRoutes(v1UserGroup)

	// init invalid routes
	initNoRoute(router)
	return router
}

func initNoRoute(router *gin.Engine) {
	router.NoRoute(handleUnknown)
}

func handleUnknown(c *gin.Context) {
	log.Error(c).Msgf("invalid URI hit: %+v", c.Request)
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Unsupported URI"})
}
