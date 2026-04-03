package action

import (
	"fmt"
	"sync/atomic"
	"time"
)

// shutdown 标志位, 用于通知所有 goroutine 停止工作并退出
var shutdown int64

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("%s work is shutting down\n", name)
			break
		}
	}
}

func Goroutines15() {
	wg.Add(2)
	go doWork("Goroutine A")
	go doWork("Goroutine B")

	// 等待 1 秒, 确保 goroutine 有机会执行
	time.Sleep(1 * time.Second)

	fmt.Println("shutdown now")

	// 通知所有 goroutine 停止工作并退出
	// StoreInt64 函数将 shutdown 标志位设置为 1,
	// 所有 goroutine 检测到标志位为 1 时, 会退出循环
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}
