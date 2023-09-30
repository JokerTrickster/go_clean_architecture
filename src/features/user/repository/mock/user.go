package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"main/common/db/mongodb"
)

type UserMock struct {
	mock.Mock
}

func (s *UserMock) InsertOneUser(ctx context.Context, dto mongodb.UserDTO) error {
	args := s.Called(dto)

	return args.Error(0)
}

func (s *UserMock) FindUser(ctx context.Context) ([]*mongodb.UserDTO, error) {
	args := s.Called()
	return args.Get(0).([]*mongodb.UserDTO), args.Error(1)
}
