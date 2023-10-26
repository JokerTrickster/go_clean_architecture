package handler

import (
	"github.com/labstack/echo/v4"
	"html"
	_interface "main/features/test/model/interface"
	"net/http"
)

type MemoryTestHandler struct {
	UseCase _interface.IMemoryTestUseCase
}

func NewMemoryTestHandler(c *echo.Echo, useCase _interface.IMemoryTestUseCase) _interface.IMemoryTestHandler {
	handler := &MemoryTestHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/test/memory", handler.MEMORY)
	return handler
}

// memory pprof 테스트
// @Router /v0.1/test/memory [get]
// @Summary memory pprof 테스트
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
func (d *MemoryTestHandler) MEMORY(c echo.Context) error {
	ctx := c.Request().Context()
	err := d.UseCase.MEMORY(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, html.EscapeString(c.Request().URL.Path))
}
