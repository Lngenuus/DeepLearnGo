package ratelimit

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCPU(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(11 * time.Millisecond)
		f := getCpuPercent()
		fmt.Println(f)
	}
}

func TestLimiter(t *testing.T) {
	l := NewDefaultLimiter()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		time.Sleep(110 * time.Millisecond)
		for j := 0; j < 300; j++ {
			time.Sleep(1 * time.Millisecond)
			go func() {
				f, err := l.Allow(context.TODO())
				if err != nil {
					fmt.Printf("%+v\n", err)
					return
				}
				time.Sleep(time.Duration(r.Intn(2000) * int(time.Millisecond)))
				f(1)
			}()
		}
	}
}
