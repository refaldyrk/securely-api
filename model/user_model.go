package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"securely-api/dto"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}

func (m User) ToDTO() dto.UserResponse {
	userResponse := dto.UserResponse{
		ID:        m.ID,
		UserID:    m.UserID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return userResponse
}
