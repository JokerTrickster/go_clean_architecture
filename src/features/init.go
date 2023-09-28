package features

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"main/common/db/mongodb"
	userHandler "main/features/user/handler"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	e.GET("/test", func(c echo.Context) error {
		cur, _ := mongodb.AdvertiseCollection.Find(context.TODO(), bson.D{{"id", 500}})
		fmt.Println(cur.RemainingBatchLength())
		return c.NoContent(http.StatusOK)
	})
	userHandler.NewUserHandler(e)
	return nil
}
