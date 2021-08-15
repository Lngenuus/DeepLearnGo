package rolling

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewWindow(t *testing.T) {
	NewDeafultWindow()
}

func TestBuckets(t *testing.T) {
	w := NewDeafultWindow()
	for key, val := range w.Buckets {
		fmt.Println("data:", key, val)
	}
}

func TestCurBucket(t *testing.T) {
	w := NewDeafultWindow()
	b := w.curBucket()
	fmt.Printf("%v\n", b)
}

func TestReduce(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w := NewDeafultWindow()
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		time.Sleep(time.Duration(r.Intn(10)) * time.Millisecond)
		for j := 0; j < 50; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(time.Duration(r.Intn(1000)) * time.Millisecond)
				w.Add(1)
			}()
		}
	}
	wg.Wait()
	// 测试查找滑动窗口的最大值
	ret := w.Reduce(func(m map[int64]*Bucket) float64 {
		max := 0.0
		for _, bucket := range m {
			max = math.Max(max, bucket.Val)
		}
		return max
	})
	fmt.Println("ret:", ret)
}
