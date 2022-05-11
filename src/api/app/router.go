package app

import (
	"EurekaHome/src/api/app/dependence"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	handlers := dependence.NewWire()
	configureMappings(router, handlers)

	return router
}

func configureMappings(router *gin.Engine, handlers dependence.HandlerContainer) {
	apiGroup := router.Group("/eureka")
	apiGroup.GET("/opportunities/description/:description/tags/:tags", handlers.GetOpportunitiesHandler.Handler)
}
