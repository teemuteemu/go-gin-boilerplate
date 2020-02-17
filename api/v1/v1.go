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
	var model models.Dummy
	errors := e.db.Find(&model, id).GetErrors()

	if len(errors) == 0 {
		c.JSON(http.StatusOK, model)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "not_found"})
}

func (e *Env) create(c *gin.Context) {
	var model models.Dummy

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if res := e.db.NewRecord(&model); res == true {
		e.db.Create(&model)
		c.JSON(http.StatusOK, gin.H{"created": model})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
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
