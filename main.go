package main

import (
	"ProjectCRUD/application/service"
	"ProjectCRUD/config"
	Models2 "ProjectCRUD/infrastructure/models"
	"ProjectCRUD/infrastructure/repositories"
	Controller2 "ProjectCRUD/presentation/Controller"
	"ProjectCRUD/presentation/router"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting Server")
	// Database
	db := config.Connect()
	fmt.Println(db)
	validator := validator.New()

	// db for user
	//
	//db.AutoMigrate(&Models.User{}, &Models.ProfileModel{})
	////
	db.Table("users").AutoMigrate(&Models2.User{})
	//db.Model(&Models.User{}).Association("profiles").Delete((Models.ProfileModel{},gorm.Cascade{}))
	db.Table("profiles").AutoMigrate(&Models2.ProfileModel{})

	//db for profile

	// Repository
	usersRepository := repositories.NewUserEpoImpl(db)
	profileRepo := repositories.NewProfileRepositoryImp(db)

	// Service
	userServices := service.NewUserServiceImp(*usersRepository, validator)

	profileService := service.NewProfileServiceImp(*profileRepo, *usersRepository, validator)
	// Controller
	userController := Controller2.NewUserController(userServices)
	fmt.Println(userController)

	profileController := Controller2.NewProfileController(profileService)
	fmt.Println(profileController)
	// Router
	routes := router.NewRouter(userController, profileController)
	//profroutes := router.NewRouterprof(profileController)

	if err := routes.Run(":1212"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start User server")
	}
}
