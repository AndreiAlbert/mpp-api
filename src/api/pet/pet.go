package pet

import (
	"gorm.io/gorm"
)

type PetCategory string

const (
	Cat PetCategory = "cat"
	Dog PetCategory = "dog"
)

type Pet struct {
	gorm.Model
	Name         string      `gorm:"size:255" json:"name"`
	Age          int         `gorm:"default:0" json:"age"`
	FavouriteToy string      `gorm:"size:255" json:"favouriteToy"`
	Category     PetCategory `gorm:"type:varchar(100)" json:"category"`
	UserId       uint        `gorm:"not null" json:"-"`
}
