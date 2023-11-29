package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"securely-api/dto"
)

type Environment struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	EnvironmentID string             `json:"environment_id" bson:"environment_id"`
	TeamID        string             `json:"team_id" bson:"team_id"`
	Key           string             `json:"key" bson:"key"`
	Value         string             `json:"value" bson:"value"`
	CreatedAt     int64              `json:"created_at" bson:"created_at"`
}

func (e Environment) ToDTO() dto.EnvironmentResponse {
	return dto.EnvironmentResponse{
		Key:   e.Key,
		Value: e.Value,
	}
}
