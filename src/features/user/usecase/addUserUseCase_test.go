package usecase

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main/common/db/mongodb"
	"main/features/user/model/interface/mocks"
	"main/features/user/model/request"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	mockAddUserRepository := new(mocks.IAddUserRepository)
	tests := []struct {
		name string
		req  request.ReqAddUser
		want error
	}{
		{"success1", request.ReqAddUser{ID: "11", Name: "ryan1", Country: "korea", Age: 12, Gender: "male"}, nil},
		{"success2", request.ReqAddUser{ID: "12", Name: "ryan2", Country: "korea", Age: 32, Gender: "male"}, nil},
		{"success3", request.ReqAddUser{ID: "13", Name: "ryan3", Country: "korea", Age: 52, Gender: "female"}, nil},
		{"failed", request.ReqAddUser{Name: "ryan3", Country: "korea", Age: 52, Gender: "female"}, fmt.Errorf("bad parameter &{ ryan3 52 korea female}")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//given
			mockAddUserRepository.On("InsertOneUser", mock.Anything, mock.MatchedBy(func(user mongodb.UserDTO) bool {
				// 여기서 필드를 무시하고 다른 필드를 일치시킬 수 있습니다.
				return user.ID == tt.req.ID &&
					user.Name == tt.req.Name &&
					user.Age == tt.req.Age &&
					user.Country == tt.req.Country &&
					user.Gender == tt.req.Gender &&
					user.IsDeleted == false
			})).Return(nil).Once()
			us := NewAddUserUseCase(mockAddUserRepository, 8*time.Second)
			//when
			got := us.Add(context.TODO(), &tt.req)
			//then
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCreateAddUserDTO(t *testing.T) {
	t.Run("ValidRequest", func(t *testing.T) {
		//given
		req := &request.ReqAddUser{
			ID:      "5",
			Name:    "ryan_dev",
			Age:     33,
			Country: "korea",
			Gender:  "male",
		}
		//when
		_, err := CreateAddUserDTO(context.Background(), req)

		//then
		assert.NoError(t, err)
	})

}
