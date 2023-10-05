package usecase

import (
	"context"
	"fmt"
	"main/common/aws/cssm"
	_interface "main/features/auth/model/interface"
	"time"
)

type GetAuthUseCase struct {
	Repository     _interface.IGetAuthRepository
	ContextTimeout time.Duration
}

func NewGetAuthUseCase(repo _interface.IGetAuthRepository, timeout time.Duration) _interface.IGetAuthUseCase {
	return &GetAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (s *GetAuthUseCase) Get(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
	defer cancel()
	url, err := cssm.AwsGetParam(ctx, "3rd_api_url")
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = HttpAPICall(ctx, url)
	fmt.Printf("usecase 로직 처리 시간 %v \n", time.Since(now))
	if err != nil {
		return err
	}

	return nil
}
