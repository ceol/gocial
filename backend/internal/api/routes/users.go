package api

import (
	"net/http"

	"github.com/ceol/gocial/internal/services"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "users",
	})
}

func detail(c *gin.Context) {
	name := c.Params.ByName("name")
	user, err := services.User.FindByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", index)
	users.GET("/:name", detail)
}
