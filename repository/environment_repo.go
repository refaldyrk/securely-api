package repository

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"securely-api/model"
)

type EnvironmentRepository struct {
	DB *qmgo.Database
}

func NewEnvironmentRepository(DB *qmgo.Database) *EnvironmentRepository {
	return &EnvironmentRepository{DB}
}

func (u *EnvironmentRepository) Find(ctx context.Context, filter bson.M) (model.Environment, error) {
	var env model.Environment

	err := u.DB.Collection("Environment").Find(ctx, filter).One(&env)
	if err != nil {
		return model.Environment{}, err
	}

	return env, nil
}

func (u *EnvironmentRepository) FindAll(ctx context.Context, filter bson.M) ([]model.Environment, error) {
	var envs []model.Environment

	err := u.DB.Collection("Environment").Find(ctx, filter).All(&envs)
	if err != nil {
		return []model.Environment{}, err
	}

	return envs, nil
}

func (u *EnvironmentRepository) Insert(ctx context.Context, data model.Environment) error {
	_, err := u.DB.Collection("Environment").InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil

}

func (u *EnvironmentRepository) Update(ctx context.Context, filter bson.M, update bson.M) error {
	err := u.DB.Collection("Environment").UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (u *EnvironmentRepository) Delete(ctx context.Context, filter bson.M) error {
	err := u.DB.Collection("Environment").Remove(ctx, filter)
	if err != nil {
		return err
	}

	return nil
	
}
