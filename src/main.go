package main

import (
	"fmt"
	"main/common/db/mongodb"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// server init setting
	if err := mongodb.Init(); err != nil {
		fmt.Println(err)
		return
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
