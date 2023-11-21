package repository

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"securely-api/model"
)

type UserRepository struct {
	DB *qmgo.Database
}

func NewUserRepository(DB *qmgo.Database) *UserRepository {
	return &UserRepository{DB}
}

func (u *UserRepository) Find(ctx context.Context, filter bson.M) (model.User, error) {
	var user model.User

	err := u.DB.Collection("User").Find(ctx, filter).One(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserRepository) FindAll(ctx context.Context, filter bson.M) ([]model.User, error) {
	var users []model.User

	err := u.DB.Collection("User").Find(ctx, filter).All(&users)
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (u *UserRepository) Insert(ctx context.Context, data model.User) error {
	_, err := u.DB.Collection("User").InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil

}

func (u *UserRepository) Update(ctx context.Context, filter bson.M, update bson.M) error {
	err := u.DB.Collection("User").UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}
