package _interface

import (
	"context"
)

type IGetAuthUseCase interface {
	Get(c context.Context) error
}
