package usecase

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"main/common/db/mongodb"
	"main/features/user/model/request"
	"main/features/user/model/response"
	"time"
)

func CreateAddUserDTO(ctx context.Context, req *request.ReqAddUser) (mongodb.UserDTO, error) {
	if req.Name == "" || req.ID == "" || req.Country == "" || req.Gender == "" || req.Age < 0 {
		return mongodb.UserDTO{}, fmt.Errorf("bad parameter %v", req)
	}
	now := time.Now()
	userDTO := mongodb.UserDTO{
		IsDeleted:   false,
		LastUpdated: &now,
		Created:     &now,
		ID:          req.ID,
		Name:        req.Name,
		Gender:      req.Gender,
		Country:     req.Country,
		Age:         req.Age,
	}
	return userDTO, nil
}

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
