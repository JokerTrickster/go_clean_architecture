package usecase

import (
	"context"
	"fmt"
	_interface "main/features/test/model/interface"
	"time"
)

type MemoryTestUseCase struct {
	Repository     _interface.ICpuTestRepository
	ContextTimeout time.Duration
}

func NewMemoryTestUseCase(repo _interface.ICpuTestRepository, timeout time.Duration) _interface.IMemoryTestUseCase {
	return &MemoryTestUseCase{Repository: repo, ContextTimeout: timeout}
}

func (s *MemoryTestUseCase) MEMORY(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()
	now := time.Now()
	fmt.Println(now)
	MemoryCalculate(ctx)
	fmt.Println("memory 함수 실행 시간 : ", time.Since(now))
	return nil
}
