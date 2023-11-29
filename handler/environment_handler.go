package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"securely-api/dto"
	"securely-api/helper"
	"securely-api/service"
)

type EnvironmentHandler struct {
	environmentService *service.EnvironmentService
}

func NewEnvironmentHandler(environmentService *service.EnvironmentService) *EnvironmentHandler {
	return &EnvironmentHandler{environmentService}
}

func (e *EnvironmentHandler) AddEnvironment(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnprocessableEntity, "unauthorized", gin.H{}))
		return
	}

	var addEnvironment dto.AddEnviroment
	err := c.ShouldBindJSON(&addEnvironment)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), gin.H{}))
		return
	}

	//Service
	environment, err := e.environmentService.AddEnvironment(c, addEnvironment, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "add environment success", environment))
	return
}
