package usecase

import (
	"context"
	"main/common/db/mongodb"
	"main/features/user/model/response"
)

func CreateGetsRes(ctx context.Context, userDTOList []mongodb.UserDTO) (response.ResGetsUser, error) {

	return response.ResGetsUser{}, nil
}
