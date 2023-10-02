package repository

import "go.mongodb.org/mongo-driver/mongo"

type GetAuthRepository struct {
	UserCollection *mongo.Collection
}
