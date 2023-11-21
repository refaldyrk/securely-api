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

func (t *TeamHandler) MyTeam(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnprocessableEntity, "unauthorized", gin.H{}))
		return
	}

	team, err := t.teamService.MyTeam(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(false, http.StatusOK, "success get my team", team))
	return

}

func (t *TeamHandler) InviteMember(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnprocessableEntity, "unauthorized", gin.H{}))
		return
	}

	teamID := c.Param("team_id")

	var teamInvReq dto.TeamInviteReq
	err := c.ShouldBindJSON(&teamInvReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), gin.H{}))
		return
	}

	err = t.teamService.InviteMember(c, userID, teamInvReq.Email, teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(false, http.StatusOK, "success invite", gin.H{}))
	return
}

func (t *TeamHandler) KickMember(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, helper.ResponseAPI(false, http.StatusUnprocessableEntity, "unauthorized", gin.H{}))
		return
	}

	teamID := c.Param("team_id")

	var teamKickReq dto.TeamKickReq
	err := c.ShouldBindJSON(&teamKickReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), gin.H{}))
		return
	}

	err = t.teamService.KickMember(c, userID, teamKickReq.Email, teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, err.Error(), gin.H{}))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(false, http.StatusOK, "success kick", gin.H{}))
	return
}
