package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) home() gin.HandlerFunc {
	return func(c *gin.Context) {
		fs, err := app.fabrics.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		pw, err := app.profilesWidth.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		ph, err := app.profilesHeight.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
			"Title":          "Home",
			"Fabrics":        fs,
			"ProfilesWidth":  pw,
			"ProfilesHeight": ph,
		})
	}
}
