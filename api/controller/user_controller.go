package controller

import (
	"main/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
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

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

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

	accessToken, err := uc.UserUsecase.CreateAccessToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}

func (uc *UserController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := uc.UserUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := uc.UserUsecase.ExtractIDFromToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := uc.UserUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
