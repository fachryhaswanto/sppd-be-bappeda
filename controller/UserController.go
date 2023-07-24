package controller

import (
	"fmt"
	"net/http"
	"os"
	"sppd/helper"
	"sppd/model"
	"sppd/request"
	"sppd/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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

func (c *userController) LoginUser(cntx *gin.Context) {
	var loginRequest request.LoginRequest

	var err = cntx.ShouldBindJSON(&loginRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := c.userService.Login(loginRequest.Username)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, "error : Kesalahan pada username atau password")

		return
	}

	var comparePassword = helper.CheckPasswordHash(user.Password, loginRequest.Password)
	if !comparePassword {
		cntx.JSON(http.StatusBadRequest, "error : Kesalahan pada username atau password")

		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uId":       user.Id,
		"uUsername": user.Username,
		"uPassword": user.Password,
		"uRole":     user.Role,

		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("APP_SECRET_KEY")))
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal generate token",
		})

		return
	}

	//send cookies back
	cntx.SetSameSite(http.SameSiteLaxMode)
	//false on secure (6th parameter) maybe want to change it to "true" when it goes online
	cntx.SetCookie("sppd-cookie", accessTokenString, 3600*5, "", "", false, true)

	user.Password = ""
	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusOK, userResponse)
}

func (c *userController) LogoutUser(cntx *gin.Context) {
	cntx.SetSameSite(http.SameSiteLaxMode)

	cntx.SetCookie("sppd-cookie", "", -1, "", "", false, true)

	cntx.JSON(http.StatusOK, "cookie deleted")
}

func (c *userController) Validate(cntx *gin.Context) {
	var user, _ = cntx.Get("user")

	var userResponse = helper.ConvertToUserResponse(user.(model.User))

	cntx.JSON(http.StatusOK, userResponse)
}
