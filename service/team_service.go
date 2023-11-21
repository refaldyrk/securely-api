package service

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"securely-api/dto"
	"securely-api/helper"
	"securely-api/model"
	"securely-api/repository"
	"time"
)

type TeamService struct {
	teamRepository *repository.TeamRepository
}

func NewTeamService(teamRepository *repository.TeamRepository) *TeamService {
	return &TeamService{teamRepository}
}

func (t *TeamService) CreateTeam(ctx context.Context, userID string, data dto.TeamReq) (dto.TeamResponse, error) {
	dataTeam := model.Team{
		ID:          primitive.NewObjectID(),
		OwnerID:     userID,
		TeamID:      uuid.NewString(),
		Name:        data.Name,
		TotalMember: 1,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	err := t.teamRepository.Insert(ctx, dataTeam)
	if err != nil {
		return dto.TeamResponse{}, err
	}

	teamMember := model.TeamMember{
		ID:        primitive.NewObjectID(),
		MemberID:  helper.NewXID(),
		UserID:    dataTeam.OwnerID,
		TeamID:    dataTeam.TeamID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	//Add To Member Collection
	err = t.teamRepository.InsertMember(ctx, teamMember)
	if err != nil {
		return dto.TeamResponse{}, err
	}

	return dataTeam.ToDTO(), nil
}
