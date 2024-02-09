package controller

import (
	"go_app/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	userList, err := userController.userService.GetUserList()
	if err != nil {
		logrus.Fatalln(err)
		context.JSON(500, err)
	}
	context.JSON(200, userList)
}
