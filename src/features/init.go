package features

import (
	"github.com/labstack/echo/v4"
	authHandler "main/features/auth/handler"
	userHandler "main/features/user/handler"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	e.GET("/test", func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	})
	userHandler.NewUserHandler(e)
	authHandler.NewAuthHandler(e)
	return nil
}
