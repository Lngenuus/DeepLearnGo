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
	for i := 0; i < 50; i++ {
		wg.Add(1)
		time.Sleep(time.Duration(r.Intn(100)) * time.Millisecond)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(r.Intn(100)) * time.Millisecond)
			w.Add(1)
		}()
	}
	wg.Wait()
	ret := w.Reduce(func(m map[int64]*Bucket) float64 {
		max := 0.0
		n := 0.0
		for _, bucket := range m {
			n++
			max = math.Max(max, bucket.val)
			fmt.Println(bucket)
		}
		return max
	})
	fmt.Println("ret:", ret)
}
