package _interface

import "context"

type ICpuTestUseCase interface {
	CPU(c context.Context) error
}
type IMemoryTestUseCase interface {
	MEMORY(c context.Context) error
}
