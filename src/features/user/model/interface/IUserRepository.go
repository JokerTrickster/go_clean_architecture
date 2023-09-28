package _interface

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"main/common/db/mongodb"
)

type IGetsUserRepository interface {
	FindUser(ctx context.Context, query bson.D) ([]mongodb.UserDTO, error)
	CountUser(ctx context.Context, query bson.D) (int32, error)
}
