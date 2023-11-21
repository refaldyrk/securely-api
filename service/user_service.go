package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"securely-api/dto"
	"securely-api/helper"
	"securely-api/model"
	"securely-api/repository"
	"time"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (u *UserService) Register(ctx context.Context, data dto.RegisterUserReq) (dto.UserResponse, error) {
	userCheck, err := u.userRepository.Find(ctx, bson.M{"email": data.Email})

	if !userCheck.ID.IsZero() || !errors.Is(err, qmgo.ErrNoSuchDocuments) {
		return dto.UserResponse{}, errors.New("email not available")
	}
	//Password Hash
	hashPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}
	userModel := model.User{
		ID:        primitive.NewObjectID(),
		UserID:    uuid.NewString(),
		Name:      data.Name,
		Email:     data.Email,
		Password:  hashPassword,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	//Insert To Database
	err = u.userRepository.Insert(ctx, userModel)
	if err != nil {
		return dto.UserResponse{}, err
	}

	//Throw Response

	return userModel.ToDTO(), nil
}

func (u *UserService) MySelf(ctx context.Context, userID string) (dto.UserResponse, error) {
	if userID == "" {
		return dto.UserResponse{}, errors.New("user id can't be empty")
	}

	user, err := u.userRepository.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return dto.UserResponse{}, err
	}

	return user.ToDTO(), nil
}

func (u *UserService) Login(ctx context.Context, req dto.LoginUserReq) (dto.UserResponse, error) {
	user, err := u.userRepository.Find(ctx, bson.M{"email": req.Email})
	if err != nil {
		return dto.UserResponse{}, errors.New("user not found")
	}

	if user.Name == "" {
		return dto.UserResponse{}, errors.New("username or password is wrong")
	}

	// compare password
	ok := helper.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		return dto.UserResponse{}, errors.New("username or password is wrong")
	}

	return user.ToDTO(), nil
}
