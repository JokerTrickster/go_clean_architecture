package _interface

import (
	"context"
	"main/features/user/model/request"
	"main/features/user/model/response"
)

type IGetsUserUseCase interface {
	Gets(c context.Context) (response.ResGetsUser, error)
}
type IAddUserUseCase interface {
	Add(c context.Context, req *request.ReqAddUser) error
}
