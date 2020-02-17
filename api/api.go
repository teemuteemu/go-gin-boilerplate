package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hello/api/v1"
	"net/http"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func ApplyRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/health", health)

	api := router.Group("/api")
	{
		apiv1.ApplyRoutes(api, db)
	}
}
