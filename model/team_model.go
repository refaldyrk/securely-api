package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"securely-api/dto"
)

type Team struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	OwnerID     string             `json:"owner_id" bson:"owner_id"`
	TeamID      string             `json:"team_id" bson:"team_id"`
	Name        string             `json:"name" bson:"name"`
	TotalMember int                `json:"total_member" bson:"total_member"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
	UpdatedAt   int64              `json:"updated_at" bson:"updated_at"`
}

func (t Team) ToDTO() dto.TeamResponse {
	teamResponse := dto.TeamResponse{
		TeamID:      t.TeamID,
		Name:        t.Name,
		TotalMember: t.TotalMember,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}

	return teamResponse
}

type TeamMember struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	MemberID  string             `json:"member_id" bson:"member_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	TeamID    string             `json:"team_id" bson:"team_id"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}
