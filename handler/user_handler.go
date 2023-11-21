package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"securely-api/dto"
	"securely-api/helper"
	"securely-api/service"
	"time"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(services *service.UserService) *UserHandler {
	return &UserHandler{service: services}
}

func (u *UserHandler) Register(c *gin.Context) {
	req := dto.RegisterUserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), gin.H{}))
		return
	}

	newUser, err := u.service.Register(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success register new user", newUser))
	return
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var loginReq dto.LoginUserReq
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), gin.H{}))
		return
	}

	user, err := u.service.Login(ctx, loginReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	// save to db

	// generate token
	token, err := helper.GenJWT(user.UserID, 24*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	ctx.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success login", gin.H{
		"token": token,
		"user":  user,
	}))
	return
}

func (u *UserHandler) MySelf(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnprocessableEntity, "unauthorized", gin.H{}))
		return
	}

	user, err := u.service.MySelf(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success retrieve data current user", user))
	return
}
