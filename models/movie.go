package models

import "gorm.io/gorm"


type Movie struct {
    gorm.Model
    Title string
    Rating int
}




