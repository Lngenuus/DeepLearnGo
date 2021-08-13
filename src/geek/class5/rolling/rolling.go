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
	WindowSec   time.Duration
	SmapleTimes int
}

type Opts func(c *Config)

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
	return time.Now().UnixNano() / 1000
}

func (w *Window) skip() int64 {
	return (time.Now().UnixNano()/1000 - w.lastAddTime) / w.BucketMilli
}

func (w *Window) curBucket() *Bucket {
	var bucket *Bucket
	var ok bool
	skip := w.skip()
	if skip > 0 {
		w.lastAddTime = skip * w.BucketMilli
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
	defer w.mu.RLock()
	w.delBucket()

	b := w.curBucket()
	b.val += val
	b.count++
	w.delBucket()
}

func NewWindow(opts ...Opts) *Window {
	cfg := defaultConfig
	for _, o := range opts {
		o(&cfg)
	}
	fmt.Println(cfg)
	return &Window{}
}
