package ratelimit

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

var (
	localcpu int64
	weight   = 0.4
)

func init() {
	go cpuproc()
}

func getCpuPercent() int64 {
	percent, _ := cpu.Percent(10*time.Millisecond, true)
	return int64(percent[0] * 100)
}

func cpuproc() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer func() {
		ticker.Stop()
		if err := recover(); err != nil {
			fmt.Println("cpuproc() error:", err)
			go cupinfo()
		}
	}()
	for range ticker.C {
		lastcpu := atomic.LoadInt64(&localcpu)
		curcpu := int64(float64(lastcpu)*weight + float64(getCpuPercent())*(1-weight))
		atomic.StoreInt64(&localcpu, curcpu)
	}
}

func cupinfo() int64 {
	res := atomic.LoadInt64(&localcpu)
	return res
}
