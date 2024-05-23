package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.Service.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := h.Service.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	loginResp := LoginResponse{
		Email: loginReq.Email,
		Jwt:   token,
	}
	c.JSON(http.StatusOK, loginResp)
}

func (h *UserHandler) GetUserWithPets(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	user, err := h.Service.GetUserAndPets(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user.Pets)
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successfull"})
}
