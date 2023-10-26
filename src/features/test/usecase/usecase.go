package usecase

import (
	"context"
	"math"
)

func CpuCalculate(ctx context.Context) {
	for i := 0; i < 10000000000; i++ {
		_ = math.Sqrt(float64(i))
	}
}

func MemoryCalculate(ctx context.Context) {
	buf := make([]byte, 1024*1024*100)
	var m map[int]string = make(map[int]string)
	for i := 0; i < 50; i++ {
		m[i] = string(buf)
	}
}
