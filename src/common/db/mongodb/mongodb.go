package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var (
	UserCollection *mongo.Collection
)

var (
	mongoHost    = "localhost" // Docker Compose에서 정의한 서비스 이름
	mongoPort    = 27026       // MongoDB 포트
	databaseName = "ryan_dev"  // 사용할 데이터베이스 이름
)

func Init() error {
	// MongoDB에 연결
	err := connectToMongoDB()
	if err != nil {
		return err
	}

	err = InitCollection()
	if err != nil {
		return err
	}
	fmt.Println("success db connect!!")
	return nil
}
func InitCollection() error {
	mongoDB := MongoClient.Database(databaseName)
	UserCollection = mongoDB.Collection("user")
	return nil
}

func connectToMongoDB() error {
	var err error
	uri := fmt.Sprintf("mongodb://%s:%d", mongoHost, mongoPort)
	clientOptions := options.Client().ApplyURI(uri)

	// MongoDB에 연결
	MongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	return nil
}
