package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterUserReq struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type LoginUserReq struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}
