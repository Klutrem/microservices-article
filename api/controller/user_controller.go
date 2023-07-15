package controller

import (
	"main/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(usecase domain.UserUsecase) UserController {
	return UserController{
		UserUsecase: usecase,
	}
}

func (uc *UserController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := uc.UserUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	c.JSON(http.StatusOK, user.ID)
}

func (uc *UserController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = uc.UserUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = uc.UserUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user.ID)
}

func (uc *UserController) GetUserId(c *gin.Context) {
	// authHeader := c.Request.Header.Get("Authorization")
	// t := strings.Split(authHeader, " ")
	// userID, err := uc.UserUsecase.ExtractIDFromToken(t[1])
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 	return
	// }
	id, ok := c.Get("user-id")
	if !ok {
		c.JSON(http.StatusForbidden, "unathorized")
	}
	c.JSON(http.StatusOK, id)

	// profile, err := uc.UserUsecase.GetProfileByID(c, userID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, userID)
}
