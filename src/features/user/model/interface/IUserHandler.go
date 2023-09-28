package _interface

import "github.com/labstack/echo/v4"

type IGetsUserHandler interface {
	Gets(c echo.Context) error
}

type IAddUserHandler interface {
	Add(c echo.Context) error
}
