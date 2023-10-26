package features

import (
	"github.com/labstack/echo/v4"
	authHandler "main/features/auth/handler"
	testHandler "main/features/test/handler"
	userHandler "main/features/user/handler"
	"net/http"
	"time"
)

func InitHandler(e *echo.Echo) error {
	e.GET("/", func(c echo.Context) error {
		time.Sleep(time.Millisecond * 50)
		return c.NoContent(http.StatusOK)
	})
	userHandler.NewUserHandler(e)
	authHandler.NewAuthHandler(e)
	testHandler.NewTestHandler(e)
	return nil
}
