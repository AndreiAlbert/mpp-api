package main

import (
	"andreialbert/mpp/src/database"
	"andreialbert/mpp/src/internal/middleware"
	"andreialbert/mpp/src/internal/pet"
	"andreialbert/mpp/src/internal/user"
	"andreialbert/mpp/src/internal/vaccination"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDb()
	db.AutoMigrate(&user.User{}, &pet.Pet{}, &vaccination.Vaccination{})

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	rPet := pet.NewPetRepository(db)
	sPet := pet.NewPetService(rPet)
	hPet := pet.NewPetHandler(sPet)

	rUser := user.NewUserRepository(db)
	sUser := user.NewUserService(rUser)
	hUser := user.NewUserHandler(sUser)

	rVacc := vaccination.NewVaccinationRepository(db)
	sVacc := vaccination.NewVaccinationService(rVacc)
	hVacc := vaccination.NewVaccinationHandler(sVacc)

	router.GET("/pets", middleware.AuthMiddleware(), hPet.AllPets)
	router.GET("/pets/:id", hPet.PetById)
	router.PUT("/pets/:id", middleware.AuthMiddleware(), hPet.UpdatePet)
	router.POST("/pets", middleware.AuthMiddleware(), hPet.CreatePet)
	router.DELETE("/pets/:id", middleware.AuthMiddleware(), hPet.DeletePet)

	router.GET("/vaccinations", hVacc.AllVaccinations)
	router.GET("/vaccinations/:id", hVacc.VaccinationById)
	router.GET("/vaccinations/cat/:id", hVacc.GetVaccinationByCatId)
	router.POST("/vaccinations", hVacc.CreateVacc)
	router.DELETE("/vaccinations/:id", hVacc.DeleteVacc)
	router.PUT("/vaccinations/:id", hVacc.UpdateVaccination)

	router.POST("/register", hUser.CreateUser)
	router.POST("/login", hUser.LoginUser)
	router.GET("/logout", hUser.Logout)

	router.Run(":8080")
}
