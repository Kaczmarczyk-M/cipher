package main

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var g string

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.GET("/tobehashed", func(c *gin.Context) {
		b := c.FullPath() == "/api/tobehashed"
		if b {
			g = c.Request.URL.Query()["inputValue"][0]
		} else {
			fmt.Println("Error")
		}
		fmt.Println(g)
	})

	r.Run(":8080")
}
