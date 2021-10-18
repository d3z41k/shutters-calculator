package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
			"title": "Home",
		})
	}
}
