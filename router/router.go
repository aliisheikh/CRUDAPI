package router

import (
	"ProjectCRUD/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(userController *Controller.UserController) *gin.Engine {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "welcome Home")
	})
	baseRouter := router.Group("api")
	userRouter := baseRouter.Group("/user")
	userRouter.POST("", userController.Create)
	userRouter.PUT(":userId", userController.Update)
	userRouter.DELETE(":userId", userController.Delete)
	userRouter.GET(":userId", userController.FindById)
	userRouter.PATCH("", userController.FindAll)
	return router

}
