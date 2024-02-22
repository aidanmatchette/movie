package main

import (
	"aidanmatchette/movie/initializers"
	"aidanmatchette/movie/internal/controllers"
	"aidanmatchette/movie/templates"

	"github.com/labstack/echo/v4"
)


func init() {
    initializers.LoadEnvVars()
    initializers.ConnectToDB()
}


func main() {


    e := echo.New()
    e.Renderer = templates.NewTemplate()

    e.GET("/movies", controllers.GetMoviePage)
    e.POST("/movies", controllers.CreateMovie)
    e.Logger.Fatal(e.Start(":8080"))
}
