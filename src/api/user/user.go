package user

import (
	"andreialbert/mpp/src/api/pet"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"column:username;uniqueIndex;not null" json:"username"`
	Password string    `gorm:"column:password;not null" json:"password"`
	Email    string    `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Pets     []pet.Pet `gorm:"foreignKey:UserId" json:"pets"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Jwt   string `json:"jwt"`
}
