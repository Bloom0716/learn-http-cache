package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (p *Provider) GetNow(c *gin.Context) {
	now := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, gin.H{
		"time": now,
	})
}

func (p *Provider) GetNowWithNoStore(c *gin.Context) {
 	c.Header("Cache-Control", "no-store")
	now := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, gin.H{
		"time": now,
	})
}
