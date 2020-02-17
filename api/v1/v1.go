package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Env struct {
	db *gorm.DB
}

type User struct {
	FirstName string `json:"first_name"`
	Age       int    `json:"age"`
}

type PostPayload struct {
	Foo string `json:"foo" binding:"required"`
}

func (e *Env) dummyGet(c *gin.Context) {
	name := c.Param("name")
	user1 := User{FirstName: name, Age: 35}

	c.JSON(http.StatusOK, user1)
}

func (e *Env) dummyPost(c *gin.Context) {
	var payload PostPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func ApplyRoutes(r *gin.RouterGroup, db *gorm.DB) {
	env := &Env{db: db}

	v1 := r.Group("/v1")
	{
		v1.GET("/dummy_get/:name", env.dummyGet)
		v1.POST("/dummy_post", env.dummyPost)
	}
}
