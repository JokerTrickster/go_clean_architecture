package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/db/mongodb"
	"main/features/auth/repository"
	"main/features/auth/usecase"
	"time"
)

func NewAuthHandler(c *echo.Echo) {
	NewGetAuthHandler(c, usecase.NewGetAuthUseCase(repository.NewGetAuthRepository(mongodb.UserCollection), 100*time.Second))
}
