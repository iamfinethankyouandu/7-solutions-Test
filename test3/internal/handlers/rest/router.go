package rest

import (
	"7solutionstest3/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(beefSummaryHandler *BeefSummaryHandler) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	v1.Use(middleware.ErrorHandler())

	beefGroup := v1.Group("/beef")
	beefGroup.GET("/summary", beefSummaryHandler.BeefSummary)

	return router
}
