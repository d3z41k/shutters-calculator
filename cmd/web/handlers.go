package main

import (
	"github.com/d3z41k/shutters-calculator/pkg/models"
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

		cs, err := app.colors.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		var options []string

		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
			"Title":    "Home",
			"Fabrics":  fs,
			"Colors":   cs,
			"Options":  options,
			"FabricID": 0,
			"ColorID":  0,
		})
	}
}

func (app *application) calc() gin.HandlerFunc {
	return func(c *gin.Context) {

		var options []string

		optionsCheck := []string{
			"outer_brackets",
			"middle_fixer",
			"profile_handle",
		}

		optionsRadio := []string{
			"set_of_guides",
			"control_stick",
			"magnetic_clasps",
		}

		for _, key := range optionsCheck {
			if c.PostForm(key) != "" {
				options = append(options, key)
			}
		}

		for _, key := range optionsRadio {
			if c.PostForm(key) != "0" {
				options = append(options, key+"_"+c.PostForm(key))
			}
		}

		fabricID, _ := strconv.Atoi(c.PostForm("fabric_id"))
		colorID, _ := strconv.Atoi(c.PostForm("color_id"))
		profileWidth, _ := strconv.ParseFloat(c.PostForm("profile_width"), 64)
		profileHeight, _ := strconv.ParseFloat(c.PostForm("profile_height"), 64)

		profileHeightRound := math.Ceil(profileHeight/10) * 10
		profileWidthRound := math.Ceil(profileWidth/10) * 10

		fs, err := app.fabrics.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		cs, err := app.colors.All()
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		fabric, err := app.fabrics.Get(fabricID)
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		profile, err := app.profiles.GetByFilter(map[string]interface{}{
			"width": int(profileWidthRound),
		})
		if err != nil {
			app.serverError(c.Writer, err)
			return
		}

		colorPrice := &models.ColorsPrices{Price: 0}

		if colorID != 1 {
			colorPrice, err = app.colorsPrices.GetByFilter(map[string]interface{}{
				"color_id": colorID,
				"width":    int(profileWidthRound),
			})
			if err != nil {
				app.serverError(c.Writer, err)
				return
			}
		}

		optionsPrice := 0

		if len(options) > 0 {
			optionsPrice, err = app.options.GetSumByKeys(options)
			if err != nil {
				app.serverError(c.Writer, err)
				return
			}
		}

		result := math.Ceil(profileHeightRound/100*profileWidthRound/100.0*float64(fabric.Category.Price)+float64(profile.Price)) + float64(colorPrice.Price) + float64(optionsPrice)

		c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
			"Title":         "Home",
			"Fabrics":       fs,
			"FabricID":      fabricID,
			"Colors":        cs,
			"ColorID":       colorID,
			"ProfileWidth":  profileWidth,
			"ProfileHeight": profileHeight,
			"Options":       options,
			"Result":        result,
		})
	}
}
