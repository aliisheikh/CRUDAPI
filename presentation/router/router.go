package router

import (
	Controller2 "ProjectCRUD/presentation/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(userController *Controller2.UserController, profileController *Controller2.ProfileController) *gin.Engine {
	router := gin.Default()

	// Define a route for the home endpoint
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome Home")
	})

	// Group routes under "/api"
	api := router.Group("/api")

	// Group user-related routes under "/api/user"
	user := api.Group("/users")
	{
		// Define CRUD routes for users
		user.POST("", userController.Create)
		user.PUT(":userId", userController.Update)
		user.DELETE(":userId", userController.Delete)
		user.GET(":userId", userController.FindById)

		// Group profile-related routes under "/:userId/profiles"
		profiles := user.Group("/:userId/profiles")
		{
			profiles.POST("", profileController.CreateP)
			profiles.PUT(":profId", profileController.UpdateP)
			profiles.DELETE(":profId", profileController.DeleteP)
			profiles.PATCH(":profId", profileController.FindByIdP)
			profiles.GET("", profileController.FindAllProfilesByUserID)
		}
		//profiles.GET("/:userId/profiles", profileController.FindAllProfilesByUserID)
	}
	return router
}
