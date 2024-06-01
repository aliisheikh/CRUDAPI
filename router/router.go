package router

import (
	"ProjectCRUD/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(userController *Controller.UserController) *gin.Engine {
	router := gin.Default()

	// Define a route for the home endpoint
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome Home")
	})

	// Group routes under "/api"
	api := router.Group("/api")

	// Group user-related routes under "/api/user"
	user := api.Group("/user")
	{
		// Define CRUD routes for users
		user.POST("", userController.Create)
		user.PUT(":userId", userController.Update)
		user.DELETE(":userId", userController.Delete)
		user.GET(":userId", userController.FindById)

	}
	return router
}

func NewRouterprof(profileController *Controller.ProfileController) *gin.Engine {
	profilerouter := gin.Default()

	// Define a route for the home endpoint
	profilerouter.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome Home")
	})

	// Group routes under "/api"
	api := profilerouter.Group("/api")

	// Group user-related routes under "/api/user"
	profile := api.Group("/user")
	{
		// Define CRUD routes for users
		profile.POST("", profileController.CreateP)
		profile.PUT(":userId", profileController.UpdateP)
		profile.DELETE(":userId", profileController.DeleteP)
		profile.GET(":userId", profileController.FindByIdP)

	}
	return profilerouter
}
