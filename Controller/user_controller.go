package Controller

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/data/request/response"
	"ProjectCRUD/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) Create(ctx *gin.Context) {
	createuserrequest := request.CreateUserReq{}
	fmt.Println(createuserrequest)
	err := ctx.ShouldBindJSON(&createuserrequest)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	userController.userService.Create(createuserrequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

	return

}

func (userController *UserController) Update(c *gin.Context) {
	updateuserrequest := request.UpdateUserReq{}
	err := c.ShouldBindJSON(&updateuserrequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userController.userService.Update(updateuserrequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)

}

func (userController *UserController) Delete(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the userService to delete the user
	err = userController.userService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (userController *UserController) FindAll(c *gin.Context) {
	usersRespose := userController.userService.FindAll()
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   usersRespose,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

//
//func (userController *UserController) FindById(context *gin.Context) {
//
//}

func (userController *UserController) FindById(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)

	}
	userController.userService.FindById(id)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}