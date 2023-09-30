package usecase

import (
	"context"
	_interface "main/features/user/model/interface"
	"main/features/user/model/request"
	"time"
)

type AddUserUseCase struct {
	Repository     _interface.IAddUserRepository
	ContextTimeout time.Duration
}

func NewAddUserUseCase(repo _interface.IAddUserRepository, timeout time.Duration) _interface.IAddUserUseCase {
	return &AddUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (s *AddUserUseCase) Add(c context.Context, req *request.ReqAddUser) error {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()
	// create update query
	userDTO, err := CreateAddUserDTO(ctx, req)
	if err != nil {
		return err
	}
	// insert one user
	err = s.Repository.InsertOneUser(ctx, userDTO)
	if err != nil {
		return err
	}
	return nil
}
