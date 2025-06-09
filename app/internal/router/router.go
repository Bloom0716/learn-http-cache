package router

import (
	"github.com/Bloom0716/learn-http-cache/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	provider := handler.NewProvider()
	router := gin.Default()

	router.GET("/health", provider.HealthCheck)

	return router
}
