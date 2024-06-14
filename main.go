package main

import (
	"ProjectCRUD/Controller"
	"ProjectCRUD/config"
	Models "ProjectCRUD/models"
	"ProjectCRUD/router"
	"ProjectCRUD/service"
	"ProjectCRUD/user_repository"
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
	db.Table("users").AutoMigrate(&Models.User{})
	//db.Model(&Models.User{}).Association("profiles").Delete((Models.ProfileModel{},gorm.Cascade{}))
	db.Table("profiles").AutoMigrate(&Models.ProfileModel{})

	//db for profile

	// Repository
	usersRepository := user_repository.NewUserEpoImpl(db)
	profileRepo := user_repository.NewProfileRepositoryImp(db)

	// Service
	userServices := service.NewUserServiceImp(*usersRepository, validator)

	profileService := service.NewProfileServiceImp(*profileRepo, *usersRepository, validator)
	// Controller
	userController := Controller.NewUserController(userServices)
	fmt.Println(userController)

	profileController := Controller.NewProfileController(profileService)
	fmt.Println(profileController)
	// Router
	routes := router.NewRouter(userController, profileController)
	//profroutes := router.NewRouterprof(profileController)

	if err := routes.Run(":1212"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start User server")
	}

}
