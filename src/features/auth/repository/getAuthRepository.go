package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	_interface "main/features/auth/model/interface"
)

func NewGetAuthRepository(userCollection *mongo.Collection) _interface.IGetAuthRepository {
	return &GetAuthRepository{userCollection}
}
