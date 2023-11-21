package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"securely-api/dto"
	"securely-api/helper"
	"securely-api/service"
)

type TeamHandler struct {
	teamService *service.TeamService
}

func NewTeamHandler(teamService *service.TeamService) *TeamHandler {
	return &TeamHandler{teamService}
}

func (t *TeamHandler) CreateTeam(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnprocessableEntity, "unauthorized", gin.H{}))
		return
	}

	var teamReq dto.TeamReq
	err := c.ShouldBindJSON(&teamReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), gin.H{}))
		return
	}

	team, err := t.teamService.CreateTeam(c, userID, teamReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(false, http.StatusOK, "success create team", team))
	return

}