package pet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	Service *PetService
}

func NewPetHandler(s *PetService) *PetHandler {
	return &PetHandler{
		Service: s,
	}
}

func (h *PetHandler) AllPets(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no user id"})
		return
	}
	pets, err := h.Service.PetsByUserId(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch pets"})
		return
	}
	c.JSON(http.StatusOK, pets)
}

func (h *PetHandler) PetById(c *gin.Context) {
	id := c.Param("id")
	pet, err := h.Service.PetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) CreatePet(c *gin.Context) {
	var newPet Pet
	if err := c.BindJSON(&newPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User id not found"})
	}
	newPet.UserId = userId.(uint)
	createdPet, err := h.Service.CreatePet(newPet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdPet)
}

func (h *PetHandler) DeletePet(c *gin.Context) {
	petId := c.Param("id")
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User id not found"})
		return
	}
	pet, err := h.Service.PetById(petId)
	if pet.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no access to this pet"})
		return
	}
	pet, err = h.Service.DeletePet(petId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) UpdatePet(c *gin.Context) {
	id := c.Param("id")
	var updatedPet Pet
	if err := c.BindJSON(&updatedPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errpr": "auth required"})
		return
	}
	existingPet, err := h.Service.PetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if existingPet.UserId != userId.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you do not have permission to update this pet"})
		return
	}
	pet, err := h.Service.UpdatePet(id, updatedPet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pet)
}
