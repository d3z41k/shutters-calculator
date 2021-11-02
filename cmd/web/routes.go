package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := gin.Default()
	r.LoadHTMLGlob("ui/html/*")
	r.Static("/static", "./ui/static")

	r.GET("/", app.home())
	r.POST("/", app.calc())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
