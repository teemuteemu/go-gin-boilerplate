package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/api"
	"hello/config"
	"hello/db"
)

func setupRouter() *gin.Engine {
	router := gin.New()

	api.ApplyRoutes(router)

	return router
}

func main() {
	config := config.GetConfig()
	db := db.Connect(&config)
	defer db.Close()

	router := setupRouter()
	port := fmt.Sprintf(":%d", config.HTTPPort)
	router.Run(port)
}
