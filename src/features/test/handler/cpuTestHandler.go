package handler

import (
	"github.com/labstack/echo/v4"
	"html"
	_interface "main/features/test/model/interface"
	"net/http"
)

type CpuTestHandler struct {
	UseCase _interface.ICpuTestUseCase
}

func NewCpuTestHandler(c *echo.Echo, useCase _interface.ICpuTestUseCase) _interface.ICpuTestHandler {
	handler := &CpuTestHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/test/cpu", handler.CPU)
	return handler
}

// cpu pprof 테스트
// @Router /v0.1/test/cpu [get]
// @Summary cpu pprof 테스트
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description
// @Description ■ errCode with 401
// @Description TOKEN_BAD : 토큰 인증 실패
// @Description POLICY_VIOLATION : 토큰 세션 정책 위반
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} int
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags test
func (d *CpuTestHandler) CPU(c echo.Context) error {
	ctx := c.Request().Context()
	err := d.UseCase.CPU(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, html.EscapeString(c.Request().URL.Path))
}
