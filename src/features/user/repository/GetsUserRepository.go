package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/db/mongodb"
	_interface "main/features/user/model/interface"
)

func NewGetsUserRepository(userCollection *mongo.Collection) _interface.IGetsUserRepository {
	return &GetsUserRepository{userCollection}
}

func (g *GetsUserRepository) FindUser(ctx context.Context) ([]mongodb.UserDTO, error) {

	return nil, nil
}
