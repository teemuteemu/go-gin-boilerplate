package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"hello/api"
	"log"
	"strconv"
)

func setupRouter() *gin.Engine {
	router := gin.New()

	api.ApplyRoutes(router)

	return router
}

func main() {
	var config Config
	err := envconfig.Process("hello", &config)

	if err != nil {
		log.Fatal(err.Error())
	}

	router := setupRouter()
	port := ":" + strconv.Itoa(config.Port)
	router.Run(port)
}
