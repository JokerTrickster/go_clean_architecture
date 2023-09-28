package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/db/mongodb"
	_interface "main/features/user/model/interface"
)

func NewGetsUserRepository(userCollection *mongo.Collection) _interface.IGetsUserRepository {
	return &GetsUserRepository{userCollection}
}

func (g *GetsUserRepository) FindUser(ctx context.Context, query bson.D) ([]mongodb.UserDTO, error) {
	cur, err := g.UserCollection.Find(ctx, query)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cur.Close(ctx)
	result := make([]mongodb.UserDTO, 0, cur.RemainingBatchLength())
	for cur.Next(ctx) {
		curUser := mongodb.UserDTO{}
		err := cur.Decode(&curUser)
		if err != nil {
			continue
		}
		result = append(result, curUser)
	}
	return result, nil
}

func (g *GetsUserRepository) CountUser(ctx context.Context, query bson.D) (int32, error) {
	count, err := g.UserCollection.CountDocuments(ctx, query)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}
