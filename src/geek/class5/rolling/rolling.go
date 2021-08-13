package rolling

import (
	"fmt"
	"sync"
	"time"
)

var defaultConfig = Config{
	WindowSec:   10,
	SmapleTimes: 100,
}

type Config struct {
	WindowSec   int64
	SmapleTimes int
}

// Window 滑动窗口的实现
type Window struct {
	Size        int
	WindowMilli int64
	Buckets     map[int64]*Bucket
	BucketMilli int64
	lastAddTime int64
	mu          *sync.RWMutex
}

type Bucket struct {
	val   float64
	count int64
}

func unixMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func (w *Window) skip() int64 {
	return (time.Now().UnixNano()/1000000 - w.lastAddTime) / w.BucketMilli
}

func (w *Window) curBucket() *Bucket {
	var bucket *Bucket
	var ok bool
	skip := w.skip()
	if skip > 0 {
		w.lastAddTime = w.lastAddTime + skip*w.BucketMilli
	}
	if bucket, ok = w.Buckets[w.lastAddTime]; !ok {
		bucket = &Bucket{}
		w.Buckets[w.lastAddTime] = bucket
	}
	return bucket
}

func (w *Window) delBucket() {
	exprid := unixMilli() - w.WindowMilli

	for tmptime := range w.Buckets {
		if tmptime < exprid {
			delete(w.Buckets, tmptime)
		}
	}
}

func (w *Window) Add(val float64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	b := w.curBucket()
	b.val += val
	b.count++
	w.delBucket()
}

func (w *Window) Reduce(f func(map[int64]*Bucket) float64) float64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return f(w.Buckets)
}

func NewWindow(cfg Config) *Window {

	windowMilli := int64(cfg.WindowSec * 1000)
	bucketMilli := windowMilli / int64(cfg.SmapleTimes)
	return &Window{
		Size:        cfg.SmapleTimes,
		WindowMilli: windowMilli,
		Buckets:     make(map[int64]*Bucket),
		BucketMilli: bucketMilli,
		lastAddTime: unixMilli(),
		mu:          &sync.RWMutex{},
	}
}

func NewDeafultWindow() *Window {
	cfg := defaultConfig
	windowMilli := int64(cfg.WindowSec * 1000)
	bucketMilli := windowMilli / int64(cfg.SmapleTimes)
	ret := &Window{
		Size:        cfg.SmapleTimes,
		WindowMilli: windowMilli,
		Buckets:     make(map[int64]*Bucket),
		BucketMilli: bucketMilli,
		lastAddTime: unixMilli(),
		mu:          &sync.RWMutex{},
	}
	fmt.Println(ret.Size, ret.WindowMilli, ret.Buckets, ret.BucketMilli, ret.lastAddTime)
	return ret
}
