package controllers

import (
	"aidanmatchette/movie/initializers"
	"aidanmatchette/movie/models"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMoviePage(c echo.Context) error {
    return c.Render(200, "index.html", "index")
}

func CreateMovie(c echo.Context) error {
    ratingStr := c.FormValue("rating")
    rating, err := strconv.ParseFloat(ratingStr, 32)

    if err != nil {
        return c.Render(402, "index", "bad rating")
    }
    
    movie := models.Movie{
    	Title:  c.FormValue("title"),
    	Rating: float32(rating),
    }

    fmt.Println(movie)
    err = initializers.DB.Create(&movie).Error

    if err != nil {
        return c.Render(402, "index", "error inputing record into DB")
    }

    return c.Render(200, "index.html", movie)
}


