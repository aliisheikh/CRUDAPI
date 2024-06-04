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
	//db.Model(&Models.User{}).Association("profiles").Delete((Models.ProfileModel{},gorm.Cascade{}))

	db.Table("users").AutoMigrate(&Models.User{})
	//db for profile
	db.Table("profile").AutoMigrate(&Models.ProfileModel{})

	// Repository
	usersRepository := user_repository.NewUserEpoImpl(db)
	profileRepo := user_repository.NewProfileRepositoryImp(db)

	// Service
	userServices := service.NewUserServiceImp(*usersRepository, validator)

	profileService := service.NewProfileServiceImp(*profileRepo, validator)
	// Controller
	userController := Controller.NewUserController(userServices)
	fmt.Println(userController)

	profileController := Controller.NewProfileController(profileService)
	fmt.Println(profileController)
	// Router
	routes := router.NewRouter(userController)
	profroutes := router.NewRouterprof(profileController)

	go func() {
		if err := routes.Run(":1212"); err != nil {
			log.Fatal().Err(err).Msg("Failed to start User server")
		}
	}()
	if err := profroutes.Run(":1213"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start profile server")
	}

}

//package main
//
//import (
//	"ProjectCRUD/Controller"
//	"ProjectCRUD/config"
//	"ProjectCRUD/models"
//	"ProjectCRUD/router"
//	"ProjectCRUD/service"
//	"ProjectCRUD/user_repository"
//	"fmt"
//	"github.com/go-playground/validator/v10"
//	"github.com/rs/zerolog/log"
//)
//
//func main() {
//	log.Info().Msg("Starting Server")
//	// Database
//	db := config.Connect()
//	fmt.Println(db)
//	validator := validator.New()
//	db.Table("users").AutoMigrate(&Models.User{})
//
//	// Repository
//	usersRepository := user_repository.NewUserEpoImpl(db)
//
//	// Service
//	userServices := service.NewUserServiceImp(*usersRepository, validator)
//
//	// Controller
//	userController := Controller.NewUserController(userServices)
//	fmt.Println(userController)
//
//	// Router
//	routes := router.NewRouter(userController)
//
//	if err := routes.Run(":1212"); err != nil {
//		{
//			return
//		}
//
//		//router := gin.Default()
//
//		// Define a route handler
//		//router.GET("/", func(c *gin.Context) {
//		//	c.JSON(http.StatusOK, gin.H{
//		//		"message": "Hello, Gin!",
//		//	})
//		//})
//
//		// Start the server
//
//	}
//}
