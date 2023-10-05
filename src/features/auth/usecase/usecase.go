package usecase

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

func HttpAPICall(ctx context.Context, url string) ([]byte, error) {
	var startTime time.Time
	trace := &httptrace.ClientTrace{
		TLSHandshakeStart: func() {
			startTime = time.Now()
		},
		TLSHandshakeDone: func(s tls.ConnectionState, err error) {
			fmt.Printf("TLS 핸드셰이크가 완료되었습니다. 경과 시간: %v\n", time.Since(startTime))
		},
	}
	traceCtx := httptrace.WithClientTrace(ctx, trace)

	req, err := http.NewRequest("GET", url, nil)
	req = req.WithContext(traceCtx) // 수정된 부분
	client := http.Client{}
	response, err := client.Do(req)
	// GET 요청 보내기
	//response, err := http.Get(url)
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
	fmt.Println(string(body))
	return body, nil
}
