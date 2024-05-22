package vaccination

import (
	"time"

	"gorm.io/gorm"
)

type Vaccination struct {
	gorm.Model
	Name  string    `gorm:"size:255;not null" json:"name"`
	Date  time.Time `gorm:"not null" json:"date"`
	PetId uint      `gorm:"not null" json:"pet_id"`
}
