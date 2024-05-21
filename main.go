package main

import (
	"ProjectCRUD/Controller"
	"ProjectCRUD/config"
	"ProjectCRUD/models"
	"ProjectCRUD/router"
	"ProjectCRUD/service"
	"ProjectCRUD/user_repository"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("Starting Server")
	// Database
	db := config.Connect()
	fmt.Println(db)
	validator := validator.New()
	db.Table("users").AutoMigrate(&Models.User{})

	// Repository
	usersRepository := user_repository.NewUserEpoImpl(db)

	// Service
	userServices := service.NewUserServiceImp(*usersRepository, validator)

	// Controller
	userController := Controller.NewUserController(userServices)
	fmt.Println(userController)

	// Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}
	if err := routes.Run(":8888"); err != nil {
		panic(err)
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
	//router := gin.Default()

	// Define a route handler
	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello, Gin!",
	//	})
	//})

	// Start the server

}
