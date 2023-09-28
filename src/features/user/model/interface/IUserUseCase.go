package _interface

import (
	"context"
	"main/features/user/model/request"
	"main/features/user/model/response"
)

type IGetsUserUseCase interface {
	Gets(c context.Context, req *request.ReqGetsUser) (response.ResGetsUser, error)
}
