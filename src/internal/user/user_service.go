package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{
		Repo: r,
	}
}

func (s *UserService) CreateUser(user User) (User, error) {
	user, err := s.Repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"user_email": user.Email,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})
	tokStr, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}
	return tokStr, nil
}

func (s *UserService) GetUserAndPets(userId uint) (User, error) {
	return s.Repo.GetUserWithPets(userId)
}
