package _interface

import "github.com/labstack/echo/v4"

type IGetAuthHandler interface {
	Get(c echo.Context) error
}
