package _interface

import (
	"context"
	"main/common/db/mongodb"
)

type IGetsUserRepository interface {
	FindUser(ctx context.Context) ([]mongodb.UserDTO, error)
}
