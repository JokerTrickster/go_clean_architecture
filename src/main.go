package main

import (
	"fmt"
	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"main/common/aws/cssm"
	"main/common/db/mongodb"
	swaggerDocs "main/docs"
	"main/features"
	"main/middleware"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // 기본 포트 번호
	}
	if err := cssm.InitAwsSsm(); err != nil {
		fmt.Println(err)
		return
	}
	e := echo.New()
	//미들웨어 초기화
	if err := middleware.InitMiddleware(e); err != nil {
		fmt.Println(err)
		return
	}
	if err := mongodb.Init(); err != nil {
		fmt.Println(err)
		return
	}
	//핸드러 초기화
	if err := features.InitHandler(e); err != nil {
		fmt.Println(err)
		return
	}
	echopprof.Wrap(e)

	// swagger 초기화
	swaggerDocs.SwaggerInfo.Host = "localhost:" + port
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.HideBanner = true
	e.Logger.Fatal(e.Start(":" + port))

}
