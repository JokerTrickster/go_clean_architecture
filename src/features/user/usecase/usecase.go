package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"main/common/db/mongodb"
	"main/features/user/model/response"
)

func CreateGetsQuery() (bson.D, error) {
	query := bson.D{{"isDeleted", false}}
	return query, nil
}

func CreateGetsRes(ctx context.Context, userDTOList []mongodb.UserDTO, count int32) (response.ResGetsUser, error) {
	userList := make([]response.GetsUserInfo, 0, len(userDTOList))
	for i := 0; i < len(userDTOList); i++ {
		curUser := response.GetsUserInfo{
			ID:      userDTOList[i].ID,
			Age:     userDTOList[i].Age,
			Country: userDTOList[i].Country,
			Gender:  userDTOList[i].Gender,
			Name:    userDTOList[i].Name,
		}
		userList = append(userList, curUser)
	}
	result := response.ResGetsUser{
		Users: userList,
		Count: count,
	}
	return result, nil
}
