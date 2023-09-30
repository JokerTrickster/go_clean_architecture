package handler

import (
	"github.com/labstack/echo/v4"
	"main/common"
	_interface "main/features/user/model/interface"
	"net/http"
)

type GetsUserHandler struct {
	UseCase _interface.IGetsUserUseCase
}

func NewGetsUserHandler(c *echo.Echo, useCase _interface.IGetsUserUseCase) _interface.IGetsUserHandler {
	handler := &GetsUserHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/user", handler.Gets)
	return handler
}

// 유저 정보 리스트 API
// @Router /v0.1/user [get]
// @Summary 유저 정보 리스트 API
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
// @Success 200 {object} response.ResGetsUser
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags user
func (d *GetsUserHandler) Gets(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := d.UseCase.Gets(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if err := common.ValidateRes(c, res); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}
