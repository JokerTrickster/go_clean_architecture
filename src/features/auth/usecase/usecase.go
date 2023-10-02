package usecase

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpAPICall(ctx context.Context, url string) ([]byte, error) {
	// GET 요청 보내기
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("GET 요청 실패: %v\n", err)
		return nil, err
	}
	defer response.Body.Close()

	// 응답 본문 읽기
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("응답 본문 읽기 실패: %v\n", err)
		return nil, err
	}
	return body, nil
}
