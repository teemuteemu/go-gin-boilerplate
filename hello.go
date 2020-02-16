package main

import (
	"github.com/gin-gonic/gin"
	"hello/api"
	"hello/config"
	"hello/db"
	"strconv"
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
	port := ":" + strconv.Itoa(config.HTTPPort)
	router.Run(port)
}
