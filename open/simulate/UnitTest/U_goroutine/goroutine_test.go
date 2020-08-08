package U_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	var c chan int
	var wg sync.WaitGroup
	const goroutineNum = 1e6

	memConsumed := func() uint64 {
		runtime.GC()
		var memStat runtime.MemStats
		runtime.ReadMemStats(&memStat)
		return memStat.Sys
	}

	release := func() {
		wg.Done()
		<-c
	}

	wg.Add(goroutineNum)
	before := memConsumed()
	for i := 0; i < goroutineNum; i++ {
		go release()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3f KB\n", float64(after-before)/goroutineNum/1000)
}
