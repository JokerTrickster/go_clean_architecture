package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/db/mongodb"
	"main/features/user/repository"
	"main/features/user/usecase"
	"time"
)

func NewUserHandler(c *echo.Echo) {
	NewGetsUserHandler(c, usecase.NewGetsUserUseCase(repository.NewGetsUserRepository(mongodb.UserCollection), 8*time.Second))
}
