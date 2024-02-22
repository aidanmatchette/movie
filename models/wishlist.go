package models

import "gorm.io/gorm"

type Wishlist struct {
    gorm.Model
    Movies []Movie
    Name string
}
