package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/qiniu/qmgo"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"securely-api/constant"
	"securely-api/dto"
	"securely-api/helper"
	"securely-api/model"
	"securely-api/repository"
	"time"
)

type EnvironmentService struct {
	environmentRepo *repository.EnvironmentRepository
	teamRepository  *repository.TeamRepository
}

func NewEnvironmentService(environmentRepository *repository.EnvironmentRepository, teamRepository *repository.TeamRepository) *EnvironmentService {
	return &EnvironmentService{environmentRepo: environmentRepository, teamRepository: teamRepository}
}

func (e *EnvironmentService) AddEnvironment(ctx context.Context, data dto.AddEnviroment, userID string) (dto.EnvironmentResponse, error) {
	//Check If User Has Access
	user, err := e.teamRepository.FindMember(ctx, bson.M{"user_id": userID, "team_id": data.TeamID})
	if err != nil {
		return dto.EnvironmentResponse{}, err
	}

	if user.Role == constant.ROLE_MEMBER {
		return dto.EnvironmentResponse{}, errors.New("u can't access")
	}

	//Check Team To Get Key
	team, err := e.teamRepository.Find(ctx, bson.M{"team_id": user.TeamID})
	if err != nil {
		return dto.EnvironmentResponse{}, err
	}

	//Check Duplicate Key
	env, err := e.environmentRepo.Find(ctx, bson.M{"key": data.Key, "team_id": team.TeamID})
	if !errors.Is(err, qmgo.ErrNoSuchDocuments) {
		return dto.EnvironmentResponse{}, errors.New("has key")
	}
	if !env.ID.IsZero() {
		return dto.EnvironmentResponse{}, errors.New("has key")
	}

	//Init Model
	environmentModel := model.Environment{
		ID:            primitive.NewObjectID(),
		EnvironmentID: uuid.NewString() + xid.New().String(),
		TeamID:        data.TeamID,
		Key:           data.Key,
		Value:         helper.EncryptAES2(data.Value, team.Teamkey),
		CreatedAt:     time.Now().Unix(),
	}

	err = e.environmentRepo.Insert(ctx, environmentModel)
	if err != nil {
		return environmentModel.ToDTO(), err
	}

	return environmentModel.ToDTO(), nil
}
