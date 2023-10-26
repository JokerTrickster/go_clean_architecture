package usecase

import (
	"context"
	"fmt"
	_interface "main/features/test/model/interface"
	"time"
)

type CpuTestUseCase struct {
	Repository     _interface.ICpuTestRepository
	ContextTimeout time.Duration
}

func NewCpuTestUseCase(repo _interface.ICpuTestRepository, timeout time.Duration) _interface.ICpuTestUseCase {
	return &CpuTestUseCase{Repository: repo, ContextTimeout: timeout}
}

func (s *CpuTestUseCase) CPU(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()
	now := time.Now()
	fmt.Println(now)
	CpuCalculate(ctx)
	fmt.Println("cpu 함수 실행 시간 : ", time.Since(now))
	return nil
}
