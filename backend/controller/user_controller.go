package controller

import (
	"go_app/service"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUserList(context *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &UserController{userService}
}

func (userController UserController) GetUserList(context *gin.Context) {
	userList := userController.userService.GetUserList()
	context.JSON(200, userList)
}
