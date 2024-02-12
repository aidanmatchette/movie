package initializers

import "aidanmatchette/movie/models"


func MigrateModels() {
    DB.AutoMigrate(&models.Movie{})
    
}
