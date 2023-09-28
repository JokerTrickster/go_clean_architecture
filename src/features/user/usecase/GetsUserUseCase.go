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
	//create query
	query, err := CreateGetsQuery()
	if err != nil {
		return response.ResGetsUser{}, err
	}
	//find user list
	userDTOList, err := s.Repository.FindUser(ctx, query)
	if err != nil {
		return response.ResGetsUser{}, err
	}
	//count user list
	count, err := s.Repository.CountUser(ctx, query)
	//create response
	res, err := CreateGetsRes(ctx, userDTOList, count)
	if err != nil {
		return response.ResGetsUser{}, err
	}
	return res, nil
}
