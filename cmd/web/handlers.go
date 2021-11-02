package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func (app *application) home() gin.HandlerFunc {
	return func(c *gin.Context) {
		fs, err := app.fabrics.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
			"Title":   "Home",
			"Fabrics": fs,
		})
	}
}

func (app *application) calc() gin.HandlerFunc {
	return func(c *gin.Context) {
		fabricID, _ := strconv.Atoi(c.PostForm("fabric_id"))
		profileWidth, _ := strconv.ParseFloat(c.PostForm("profile_width"), 64)
		profileHeight, _ := strconv.ParseFloat(c.PostForm("profile_height"), 64)

		profileHeightRound := math.Ceil(profileHeight/10) * 10
		profileWidthRound := math.Ceil(profileWidth/10) * 10

		fs, err := app.fabrics.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		fabric, err := app.fabrics.Get(fabricID)
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		profile, err := app.profiles.GetByField("width", int(profileWidthRound))
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		result := math.Ceil(profileHeightRound/100*profileWidthRound/100.0*float64(fabric.Category.Price) + float64(profile.Price))

		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
			"Title":         "Home",
			"Fabrics":       fs,
			"ProfileWidth":  profileWidth,
			"ProfileHeight": profileHeight,
			"Result":        result,
		})
	}
}
