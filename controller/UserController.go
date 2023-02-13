package controller

import (
	"net/http"
	"sppd/helper"
	"sppd/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (c *userController) GetUser(cntx *gin.Context) {
	var username = cntx.Param("username")
	var password = cntx.Param("password")

	var user, err = c.userService.FindByUser(username, password)
	if err != nil {
		cntx.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}
