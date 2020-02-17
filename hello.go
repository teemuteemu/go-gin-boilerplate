package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hello/api"
	"hello/config"
	"hello/db"
	"hello/middleware"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	api.ApplyRoutes(router, db)
	middleware.ApplyMiddlewares(router)

	return router
}

func main() {
	config := config.GetConfig()
	db := db.Connect(&config)
	defer db.Close()

	router := setupRouter(db)

	port := fmt.Sprintf(":%d", config.HTTPPort)
	router.Run(port)
}
