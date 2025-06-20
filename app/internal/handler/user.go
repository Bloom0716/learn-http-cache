package handler

import (
	"net/http"
	"strconv"

	"github.com/Bloom0716/learn-http-cache/internal/model"
	"github.com/gin-gonic/gin"
)

func (p *Provider) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	user, err := p.Repository.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if match := c.GetHeader("If-None-Match"); match != "" && match == user.Etag {
		c.Status(http.StatusNotModified)
		return
	}

	c.Header("Cache-Control", "no-cache")
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":   user.ID,
			"name": user.Name,
		},
		"etag": user.Etag,
	})
}

type CreateUserRequestParams struct {
	Name string `json:"name" binding:"required"`
}

func (p *Provider) CreateUser(c *gin.Context) {
	var params CreateUserRequestParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user := &model.User{Name: params.Name}
	createdUser, err := p.Repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.Header("ETag", createdUser.Etag)
	c.Header("Cache-Control", "no-store")
	c.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":   createdUser.ID,
			"name": createdUser.Name,
		},
		"etag": createdUser.Etag,
	})
}
