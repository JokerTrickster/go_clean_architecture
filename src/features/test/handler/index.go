package handler

import (
	"github.com/labstack/echo/v4"
	"main/features/test/repository"
	"main/features/test/usecase"
	"time"
)

func NewTestHandler(c *echo.Echo) {
	NewCpuTestHandler(c, usecase.NewCpuTestUseCase(repository.NewCpuTestRepository(), 100*time.Second))
	NewMemoryTestHandler(c, usecase.NewMemoryTestUseCase(repository.NewMemoryTestRepository(), 100*time.Second))
}
