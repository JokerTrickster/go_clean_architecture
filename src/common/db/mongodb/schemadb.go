package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AdvertiseDTO struct {
	MID           primitive.ObjectID `bson:"_id,omitempty"`
	ID            int64              `bson:"id"`
	Name          string             `bson:"name,omitempty"`
	ImageUrl      string             `bson:"image_url,omitempty"`
	LandingUrl    string             `bson:"landing_url,omitempty"`
	Weight        int64              `bson:"weight"`
	TargetCountry string             `bson:"target_country"`
	TargetGender  string             `bson:"target_gender"`
	Reward        int32              `bson:"reward"`
	IsDeleted     bool               `bson:"isDeleted,omitempty"`
	LastUpdated   time.Time          `bson:"lastUpdated,omitempty"`
}
type UserDTO struct {
	MID         primitive.ObjectID `bson:"_id,omitempty"`
	IsDeleted   bool               `bson:"isDeleted"`
	Created     time.Time          `bson:"created"`
	LastUpdated time.Time          `bson:"lastUpdated"`
	ID          string             `bson:"id"`
	Name        string             `bson:"name"`
	Gender      string             `bson:"gender"`
	Country     string             `bson:"country"`
	Age         int32              `bson:"age"`
}
type AdvertiseSendHistoryDTO struct {
	MID         primitive.ObjectID `bson:"_id,omitempty"`
	IsDeleted   bool               `bson:"isDeleted"`
	Created     time.Time          `bson:"created"`
	LastUpdated time.Time          `bson:"lastUpdated"`
	UserID      string             `bson:"userID"`
	Advertises  []Advertise        `bson:"advertises"`
}
type Advertise struct {
	AdvertiseID int32 `bson:"advertiseID"`
	IsDeposit   bool  `bson:"isDeposit"`
}

type RewardDTO struct {
	MID         primitive.ObjectID `bson:"_id,omitempty"`
	IsDeleted   bool               `bson:"isDeleted"`
	Created     time.Time          `bson:"created"`
	LastUpdated time.Time          `bson:"lastUpdated"`
	UserID      string             `bson:"userID"`
	Point       int32              `bson:"point"`
}
type RewardHistoryDTO struct {
	MID             primitive.ObjectID `bson:"_id,omitempty"`
	IsDeleted       bool               `bson:"isDeleted"`
	Created         time.Time          `bson:"created"`
	LastUpdated     time.Time          `bson:"lastUpdated"`
	UserID          string             `bson:"userID"`
	LastAdvertiseID int32              `bson:"lastAdvertiseID"`
	Type            string             `bson:"type"`
	CurrentPoint    int32              `bson:"currentPoint"`
	PreviousPoint   int32              `bson:"previousPoint"`
}
