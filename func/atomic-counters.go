package _func

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounters() {
	// 核心作用：原子操作，确保多个 Goroutine 对共享变量的访问是安全的
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				ops.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops.Load())
}
