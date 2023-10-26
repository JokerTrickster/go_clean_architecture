package _interface

import "github.com/labstack/echo/v4"

type ICpuTestHandler interface {
	CPU(c echo.Context) error
}
type IMemoryTestHandler interface {
	MEMORY(c echo.Context) error
}
