package rolling

import (
	"sync"
	"time"
)

var defaultWindowSec = 10
var defaultSampleTimes = 100

// Window 滑动窗口的实现
type Window struct {
	Size        int
	WindowMilli int64
	Buckets     map[int64]*Bucket
	BucketMilli int64
	lastAddTime int64
	mu          *sync.RWMutex
}

// Bucket 储存数据桶
type Bucket struct {
	Val   float64
	Count int64
}

// 毫秒时间单位生成函数
func unixMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

// skip 当前走过的时间小窗
func (w *Window) skip() int64 {
	return (time.Now().UnixNano()/1000000 - w.lastAddTime) / w.BucketMilli
}

// curBucket
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

// 遍历map 找到所有时间节点小于当前窗口的
// delBucket 删除过期的小窗
func (w *Window) delBucket() {
	exprid := unixMilli() - w.WindowMilli

	for tmptime := range w.Buckets {
		if tmptime < exprid {
			delete(w.Buckets, tmptime)
		}
	}
}

// Add 添加一个指标数据
func (w *Window) Add(val float64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	b := w.curBucket()
	b.Val += val
	b.Count++
	w.delBucket()
}

// 传入操作函数处理数据
// Reduce 聚合函数
func (w *Window) Reduce(f func(map[int64]*Bucket) float64) float64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return f(w.Buckets)
}

// NewWindow 窗口生成可以传入窗口大小和采样频次
func NewWindow(windowSec int, sampleTimes int) *Window {

	windowMilli := int64(windowSec * 1000)
	bucketMilli := windowMilli / int64(sampleTimes)
	return &Window{
		Size:        sampleTimes,
		WindowMilli: windowMilli,
		Buckets:     make(map[int64]*Bucket),
		BucketMilli: bucketMilli,
		lastAddTime: unixMilli(),
		mu:          &sync.RWMutex{},
	}
}

// 默认窗口生成
func NewDeafultWindow() *Window {
	windowMilli := int64(defaultWindowSec * 1000)
	bucketMilli := windowMilli / int64(defaultSampleTimes)
	return &Window{
		Size:        defaultSampleTimes,
		WindowMilli: windowMilli,
		Buckets:     make(map[int64]*Bucket),
		BucketMilli: bucketMilli,
		lastAddTime: unixMilli(),
		mu:          &sync.RWMutex{},
	}
}
