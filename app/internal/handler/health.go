package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Provider) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
