package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ibnu-a/Golang-RESTful-API/dto"
	"github.com/ibnu-a/Golang-RESTful-API/helper"
	"github.com/ibnu-a/Golang-RESTful-API/service"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWtService
}

func NewUserController(userService service.UserService, jwtService service.JWtService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(contex *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := contex.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		contex.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := contex.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	contex.JSON(http.StatusOK, res)
}

func (c *userController) Profile(contex *gin.Context) {
	authHeader := contex.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", user)
	contex.JSON(http.StatusOK, res)
}
