package repository

import "go.mongodb.org/mongo-driver/mongo"

type GetsUserRepository struct {
	UserCollection *mongo.Collection
}
