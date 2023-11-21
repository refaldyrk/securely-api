package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"securely-api/constant"
	"securely-api/dto"
	"securely-api/helper"
	"securely-api/model"
	"securely-api/repository"
	"time"
)

type TeamService struct {
	teamRepository *repository.TeamRepository
	userRepository *repository.UserRepository
}

func NewTeamService(teamRepository *repository.TeamRepository, userRepository *repository.UserRepository) *TeamService {
	return &TeamService{teamRepository: teamRepository, userRepository: userRepository}
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

	accessKey := helper.NewXID()

	teamMember := model.TeamMember{
		ID:        primitive.NewObjectID(),
		MemberID:  helper.NewXID(),
		UserID:    dataTeam.OwnerID,
		TeamID:    dataTeam.TeamID,
		Role:      constant.ROLE_OWNER,
		AccessKey: helper.EncryptAES(accessKey),
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

func (t *TeamService) MyTeam(ctx context.Context, userID string) ([]model.TeamMember, error) {
	teamByUserID, err := t.teamRepository.GetTeamByUserID(ctx, userID)
	if err != nil {
		return []model.TeamMember{}, err
	}

	return teamByUserID, nil
}

func (t *TeamService) InviteMember(ctx context.Context, userID, email, teamID string) error {
	user, err := t.userRepository.Find(ctx, bson.M{"email": email})
	if err != nil || user.ID.IsZero() {
		if errors.Is(err, qmgo.ErrNoSuchDocuments) {
			return errors.New("user not found")
		}
		return err
	}

	//Check Role
	inviter, err := t.teamRepository.FindMember(ctx, bson.M{"team_id": teamID, "user_id": userID})
	if err != nil {
		if errors.Is(err, qmgo.ErrNoSuchDocuments) {
			return errors.New(" not found")
		}
		return err
	}

	if inviter.Role != constant.ROLE_OWNER {
		return errors.New("access denied bcs u not owner, pls contact your owner to invite someone")
	}

	//Find Team
	team, err := t.teamRepository.Find(ctx, bson.M{"team_id": teamID})
	if err != nil {
		return err
	}

	//Validate If Invite Is Exists
	_, err = t.teamRepository.FindMember(ctx, bson.M{"user_id": user.UserID})
	if !errors.Is(err, qmgo.ErrNoSuchDocuments) {
		return errors.New(" available")
	}

	//Add To Member
	accessKey := helper.NewXID()
	teamMember := model.TeamMember{
		ID:        primitive.NewObjectID(),
		MemberID:  helper.NewXID(),
		UserID:    user.UserID,
		TeamID:    teamID,
		Role:      constant.ROLE_MEMBER,
		AccessKey: helper.EncryptAES(accessKey),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err = t.teamRepository.InsertMember(ctx, teamMember)
	if err != nil {
		return err
	}

	//increment Your Team
	err = t.teamRepository.Update(ctx, bson.M{"team_id": teamID}, bson.M{"updated_at": time.Now().Unix(), "total_member": team.TotalMember + 1})
	if err != nil {
		return err
	}

	return nil
}
