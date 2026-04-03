package action

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

// 为了避免竞态条件, 可以使用原子函数，atomic包
// atomic 包提供了原子操作函数，用于在并发环境中安全地更新共享变量

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 原子加载 counter 的当前值
		// 强制只有一个 goroutine 可以执行 AddInt64 操作
		// 每当执行 AddInt64 操作时, goroutine 会同步更新 counter 的值
		atomic.AddInt64(&counter, 1)

		// 退出当前 goroutine 并允许其他 goroutine 执行
		runtime.Gosched()
	}
}

func Goroutines13() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Counter:", counter)
}
