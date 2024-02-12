package main

import (
	"aidanmatchette/movie/initializers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
    initializers.LoadEnvVars()
    initializers.ConnectToDB()
}


func main() {


    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, world")
    })
    e.Logger.Fatal(e.Start(":8080"))
}
