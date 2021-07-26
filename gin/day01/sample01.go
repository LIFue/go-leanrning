package day01

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOneRoute() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := route.Run(":8080")
	if err != nil {
		return 
	}
}
