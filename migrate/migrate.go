package main

import (
	"aidanmatchette/movie/initializers"
	"aidanmatchette/movie/models"
)

func init() {
    initializers.LoadEnvVars()
    initializers.ConnectToDB()
}
    

func main() {
    initializers.DB.AutoMigrate(&models.Movie{})
}
