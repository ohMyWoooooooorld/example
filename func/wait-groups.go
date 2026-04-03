package _func

import (
	"fmt"
	"sync"
	"time"
)

func worker3(id int) {
	fmt.Printf("worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	// 核心改造点：手动调用 Add(1) 并使用标准的 go 关键字
	for i := 1; i <= 5; i++ {
		// 1. 每次启动 Goroutine 前，告知 WaitGroup 计数器加 1
		wg.Add(1)

		// 2. 启动 Goroutine
		go func(workerID int) { // 传入 i 的值以避免闭包陷阱
			// 3. 确保 Goroutine 结束时调用 Done()
			defer wg.Done()

			worker3(workerID)
		}(i) // 立即执行函数，传入当前的 i 值
	}

	wg.Wait()
}
