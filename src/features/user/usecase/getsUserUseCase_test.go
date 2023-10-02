package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main/common/db/mongodb"
	"main/features/user/model/interface/mocks"
	"main/features/user/model/response"
	"testing"
	"time"
)

func TestGets(t *testing.T) {
	mockGetsUserRepository := new(mocks.IGetsUserRepository)
	tests := []struct {
		name string
		want response.ResGetsUser
	}{
		{"success1", response.ResGetsUser{Users: []response.GetsUserInfo{response.GetsUserInfo{ID: "1", Age: 33, Country: "korea", Name: "ryan", Gender: "male"}}, Count: 1}},
		{"success2", response.ResGetsUser{Users: []response.GetsUserInfo{response.GetsUserInfo{ID: "1", Age: 33, Country: "korea", Name: "ryan", Gender: "male"}}, Count: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//given
			query, err := CreateGetsQuery()
			assert.NoError(t, err)
			now := time.Now()
			mockGetsUserRepository.On("FindUser", mock.Anything, query).Return([]mongodb.UserDTO{
				{ID: "1", IsDeleted: false, Created: &now, LastUpdated: &now, Name: "ryan", Gender: "male", Country: "korea", Age: 33},
			}, nil).Once()

			mockGetsUserRepository.On("CountUser", mock.Anything, query).Return(int32(1), nil).Once()
			us := NewGetsUserUseCase(mockGetsUserRepository, 8*time.Second)

			//when
			got, err := us.Gets(context.TODO())
			//then
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCreateGetsRes(t *testing.T) {
	// 테스트에 사용할 가상의 데이터 생성
	userDTOList := []mongodb.UserDTO{
		{
			ID:      "1",
			Age:     25,
			Country: "USA",
			Gender:  "Male",
			Name:    "John",
		},
		{
			ID:      "2",
			Age:     30,
			Country: "Canada",
			Gender:  "Female",
			Name:    "Alice",
		},
	}

	count := int32(len(userDTOList))

	// CreateGetsRes 함수 호출
	result, err := CreateGetsRes(context.Background(), userDTOList, count)

	// 에러가 없는지 확인
	assert.NoError(t, err)

	// 반환된 결과 확인
	assert.Equal(t, count, result.Count)
	assert.Len(t, result.Users, len(userDTOList))

	// 개별 사용자 정보 확인
	expectedUsers := []response.GetsUserInfo{
		{
			ID:      "1",
			Age:     25,
			Country: "USA",
			Gender:  "Male",
			Name:    "John",
		},
		{
			ID:      "2",
			Age:     30,
			Country: "Canada",
			Gender:  "Female",
			Name:    "Alice",
		},
	}

	for i, expectedUser := range expectedUsers {
		assert.Equal(t, expectedUser, result.Users[i])
	}
}
