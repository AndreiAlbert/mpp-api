package vaccination

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type VaccinationHandler struct {
	Service *VaccinationService
}

func NewVaccinationHandler(s *VaccinationService) *VaccinationHandler {
	return &VaccinationHandler{
		Service: s,
	}
}

func (h *VaccinationHandler) AllVaccinations(c *gin.Context) {
	vaccs, err := h.Service.AllVaccinations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch vaccs"})
		return
	}
	c.JSON(http.StatusOK, vaccs)
}

func (h *VaccinationHandler) CreateVacc(c *gin.Context) {
	var input struct {
		Name  string `json:"name"`
		Date  string `json:"date"`
		PetId uint   `json:"pet_id"`
	}
	var vacc Vaccination
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	vacc.Date = date
	vacc.PetId = input.PetId
	vacc.Name = input.Name
	newVacc, err := h.Service.CreateVaccination(vacc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newVacc)
}

func (h *VaccinationHandler) GetVaccinationByCatId(c *gin.Context) {
	petId := c.Param("id")
	vaccs, err := h.Service.VaccinationsByCatId(petId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(petId)
	c.JSON(http.StatusOK, vaccs)
}

func (h *VaccinationHandler) DeleteVacc(c *gin.Context) {
	vaccId := c.Param("id")
	vacc, err := h.Service.DeleteVaccination(vaccId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vacc)
}

func (h *VaccinationHandler) VaccinationById(c *gin.Context) {
	vaccId := c.Param("id")
	vacc, err := h.Service.VaccinationById(vaccId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vacc)
}

func (h *VaccinationHandler) UpdateVaccination(c *gin.Context) {
	var input struct {
		Name  string `json:"name"`
		Date  string `json:"date"`
		PetId uint   `json:"pet_id"`
	}
	var vacc Vaccination
	id := c.Param("id")
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	vacc.Date = date
	vacc.PetId = input.PetId
	vacc.Name = input.Name
	newVacc, err := h.Service.UpdateVaccination(id, vacc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newVacc)

}
