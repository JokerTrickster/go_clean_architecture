package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/db/mongodb"
	_interface "main/features/user/model/interface"
)

func NewAddUserRepository(userCollection *mongo.Collection) _interface.IAddUserRepository {
	return &AddUserRepository{userCollection}
}

func (g *AddUserRepository) InsertOneUser(ctx context.Context, userDTO mongodb.UserDTO) error {
	_, err := g.UserCollection.InsertOne(ctx, userDTO)
	if err != nil {
		return err
	}
	return nil
}
