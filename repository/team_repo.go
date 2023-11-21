package repository

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"securely-api/model"
)

type TeamRepository struct {
	DB *qmgo.Database
}

func NewTeamRepository(DB *qmgo.Database) *TeamRepository {
	return &TeamRepository{DB}
}

func (t *TeamRepository) Find(ctx context.Context, filter bson.M) (model.Team, error) {
	var team model.Team

	err := t.DB.Collection("Team").Find(ctx, filter).One(&team)
	if err != nil {
		return model.Team{}, err
	}

	return team, nil
}

func (t *TeamRepository) FindAll(ctx context.Context, filter bson.M) ([]model.Team, error) {
	var teams []model.Team

	err := t.DB.Collection("Team").Find(ctx, filter).All(&teams)
	if err != nil {
		return []model.Team{}, err
	}

	return teams, nil
}

func (t *TeamRepository) Insert(ctx context.Context, data model.Team) error {
	_, err := t.DB.Collection("Team").InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil

}

func (t *TeamRepository) Update(ctx context.Context, filter bson.M, update bson.M) error {
	err := t.DB.Collection("Team").UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (t *TeamRepository) Delete(ctx context.Context, filter bson.M) error {
	err := t.DB.Collection("Team").Remove(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (t *TeamRepository) FindMember(ctx context.Context, filter bson.M) (model.TeamMember, error) {
	var member model.TeamMember

	err := t.DB.Collection("Member").Find(ctx, filter).One(&member)
	if err != nil {
		return model.TeamMember{}, err
	}

	return member, nil
}

func (t *TeamRepository) FindAllMember(ctx context.Context, filter bson.M) ([]model.TeamMember, error) {
	var members []model.TeamMember

	err := t.DB.Collection("Member").Find(ctx, filter).All(&members)
	if err != nil {
		return []model.TeamMember{}, err
	}

	return members, nil
}

func (t *TeamRepository) InsertMember(ctx context.Context, data model.TeamMember) error {
	_, err := t.DB.Collection("Member").InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil

}

func (t *TeamRepository) UpdateMember(ctx context.Context, filter bson.M, update bson.M) error {
	err := t.DB.Collection("Member").UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (t *TeamRepository) DeleteMember(ctx context.Context, filter bson.M) error {
	err := t.DB.Collection("Member").Remove(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
