package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserWithPets(userId uint) (User, error) {
	var usr User
	if err := r.db.Preload("Pets").First(&usr, userId).Error; err != nil {
		return User{}, err
	}
	return usr, nil
}

func (r *UserRepository) GetUserByEmail(email string) (User, error) {
	var user User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user User) (User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	user.Password = string(hashedPass)
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
