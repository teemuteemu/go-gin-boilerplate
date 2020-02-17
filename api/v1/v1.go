package apiv1

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hello/models"
	"net/http"
)

type Env struct {
	db *gorm.DB
}

/*
type PostPayload struct {
	SomeField string `json:"" binding:"required"`
}
*/

func (e *Env) list(c *gin.Context) {
	var dummies []models.Dummy
	e.db.Find(&dummies)

	c.JSON(http.StatusOK, dummies)
}

func (e *Env) get(c *gin.Context) {
	id := c.Param("id")
	var dummy models.Dummy
	errors := e.db.Find(&dummy, id).GetErrors()

	if len(errors) == 0 {
		c.JSON(http.StatusOK, dummy)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{})
}

func (e *Env) create(c *gin.Context) {
	var dummy models.Dummy
	c.BindJSON(&dummy)

	if res := e.db.NewRecord(&dummy); res == true {
		e.db.Create(&dummy)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
}

func ApplyRoutes(r *gin.RouterGroup, db *gorm.DB) {
	env := &Env{db: db}

	v1 := r.Group("/v1")
	{
		v1.GET("/dummy", env.list)
		v1.GET("/dummy/:id", env.get)
		v1.POST("/dummy", env.create)
	}
}
