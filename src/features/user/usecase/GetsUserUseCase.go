package usecase

import (
	"context"
	_interface "main/features/user/model/interface"
	"main/features/user/model/request"
	"main/features/user/model/response"
	"time"
)

type GetsUserUseCase struct {
	Repository     _interface.IGetsUserRepository
	ContextTimeout time.Duration
}

func NewGetsUserUseCase(repo _interface.IGetsUserRepository, timeout time.Duration) _interface.IGetsUserUseCase {
	return &GetsUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (s *GetsUserUseCase) Gets(c context.Context, req *request.ReqGetsUser) (response.ResGetsUser, error) {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()
	userDTOList, err := s.Repository.FindUser(ctx)
	if err != nil {
		return response.ResGetsUser{}, err
	}
	res, err := CreateGetsRes(ctx, userDTOList)
	if err != nil {
		return response.ResGetsUser{}, err
	}
	return res, nil
}
