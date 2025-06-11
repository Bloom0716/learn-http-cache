package router

import (
	"github.com/Bloom0716/learn-http-cache/internal/database"
	"github.com/Bloom0716/learn-http-cache/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(repository database.Repository) *gin.Engine {
	provider := handler.NewProvider(repository)
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/health", provider.HealthCheck)
		api.GET("/now", provider.GetNow)
		api.GET("/now/no-store", provider.GetNowWithNoStore)
	}
	user := api.Group("/users")
	{
		user.GET("/:id", provider.GetUserByID)
		user.POST("/create", provider.CreateUser)
	}

	return router
}
