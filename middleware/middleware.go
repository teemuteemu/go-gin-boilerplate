package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApplyMiddlewares(router *gin.Engine) {
	router.Use(cors.Default())
}
