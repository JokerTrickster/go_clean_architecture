package handler

import (
	"github.com/labstack/echo/v4"
	_interface "main/features/auth/model/interface"
	"net/http"
)

type GetAuthHandler struct {
	UseCase _interface.IGetAuthUseCase
}

func NewGetAuthHandler(c *echo.Echo, useCase _interface.IGetAuthUseCase) _interface.IGetAuthHandler {
	handler := &GetAuthHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/auth/accessToken", handler.Get)
	return handler
}

// 액세스 토큰 생성하기
// @Router /v0.1/auth/accessToken [get]
// @Summary 액세스 토큰 생성하기
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
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags auth
func (d *GetAuthHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	err := d.UseCase.Get(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, true)
}
