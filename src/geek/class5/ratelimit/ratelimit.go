package ratelimit

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"github.com/dogslee/deep-learn-go/src/geek/class5/rolling"
)

var initTime = time.Now()

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type Config struct {
	WindowSec    int
	SampleTimes  int
	CpuCheckStop int64
}

type RateLimit interface {
	Allow(ctx context.Context) (func(ok int), error)
}

type Limiter struct {
	cpu          func() int64
	cpuCheckStop int64
	inFlight     int64
	windowSec    int
	sampleTimes  int
	sampleWidth  time.Duration
	passStat     rolling.Window
	rtStat       rolling.Window
	stopTime     atomic.Value
	maxPassCache atomic.Value
	minRtCache   atomic.Value
}

type TmpCache struct {
	val  int64
	time time.Time
}

func (l *Limiter) Allow(ctx context.Context) (func(ok int), error) {
	if l.checkStop() {
		return nil, errors.New("流量被限制")
	}
	atomic.AddInt64(&l.inFlight, 1)
	stime := time.Since(initTime)
	return func(ok int) {
		rt := (time.Since(initTime) - stime) / time.Millisecond
		l.rtStat.Add(float64(rt))
		if ok == 1 {
			l.passStat.Add(1)
		}
		atomic.AddInt64(&l.inFlight, -1)
	}, nil
}

func NewDefaultLimiter() *Limiter {
	return &Limiter{
		cpu:          cupinfo,
		windowSec:    10,
		sampleTimes:  100,
		cpuCheckStop: 8000,
		sampleWidth:  time.Duration(10) * time.Second / time.Duration(100),
		passStat:     *rolling.NewDeafultWindow(),
		rtStat:       *rolling.NewDeafultWindow(),
	}
}

func NewLimiter(c *Config) *Limiter {
	return &Limiter{
		cpu:          cupinfo,
		windowSec:    c.WindowSec,
		sampleTimes:  c.SampleTimes,
		cpuCheckStop: c.CpuCheckStop,
		sampleWidth:  time.Duration(c.WindowSec) * time.Second / time.Duration(c.SampleTimes),
		passStat:     *rolling.NewWindow(c.WindowSec, c.SampleTimes),
		rtStat:       *rolling.NewWindow(c.WindowSec, c.SampleTimes),
	}
}

func (l *Limiter) timespan(lastTime time.Time) int {
	runtime := time.Since(lastTime)
	return int(runtime / l.sampleWidth)
}

func (l *Limiter) maxInFlight() int64 {
	res := l.maxPass() * l.minRt() / (l.sampleWidth.Milliseconds())
	fmt.Println("maxFlight:", res)
	return res
}

func (l *Limiter) maxPass() int64 {
	passCache := l.maxPassCache.Load()
	if passCache != nil {
		ps := passCache.(*TmpCache)
		if l.timespan(ps.time) < 1 {
			return int64(ps.val)
		}
	}
	tmpMaxPass := int64(l.passStat.Reduce(func(m map[int64]*rolling.Bucket) float64 {
		max := 0.0
		for _, bucket := range m {
			max = math.Max(max, bucket.Val)
		}
		fmt.Println("max:", max)
		return max
	}))
	l.maxPassCache.Store(&TmpCache{
		val:  tmpMaxPass,
		time: time.Now(),
	})
	return tmpMaxPass
}

func (l *Limiter) minRt() int64 {
	rtCache := l.minRtCache.Load()
	if rtCache != nil {
		rt := rtCache.(*TmpCache)
		if l.timespan(rt.time) < 1 {
			return int64(rt.val)
		}
	}
	tmpMinRt := int64(l.rtStat.Reduce((func(m map[int64]*rolling.Bucket) float64 {
		min := float64(INT_MAX)
		for _, bucket := range m {
			if bucket.Count > 0 {
				min = math.Min(bucket.Val/float64(bucket.Count), min)
			}
		}
		fmt.Println("min:", min)
		return min
	})))
	l.maxPassCache.Store(&TmpCache{
		val:  tmpMinRt,
		time: time.Now(),
	})
	return tmpMinRt
}

func (l *Limiter) checkStop() bool {
	cput := l.cpu()
	if cput < l.cpuCheckStop {
		lastStopTime, _ := l.stopTime.Load().(time.Duration)
		if lastStopTime == 0 {
			return false
		}
		if time.Since(initTime)-lastStopTime <= time.Second {
			inFlight := atomic.LoadInt64(&l.inFlight)
			return inFlight > 1 && inFlight > l.maxInFlight()
		}
		l.stopTime.Store(time.Duration(0))
		return false
	}
	inFlight := atomic.LoadInt64(&l.inFlight)
	fmt.Println("cupt:", cput, "inFlight:", inFlight)
	stop := inFlight > 1 && inFlight > l.maxInFlight()
	if stop {
		lastStopTime, _ := l.stopTime.Load().(time.Duration)
		if lastStopTime != 0 {
			return stop
		}
		l.stopTime.Store(time.Since(initTime))
	}
	return stop
}
