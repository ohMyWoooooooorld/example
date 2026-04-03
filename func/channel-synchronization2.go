package _func

import (
	"fmt"
	"time"
)

func worker(id int, limiter chan struct{}) {
	// 任务完成后的操作
	defer func() {
		// 从 limiter 通道中释放一个空结构体，允许其他 goroutine 继续执行
		<-limiter
		wg.Done()
	}()

	// 任务处理开始
	// 获取一个任务，并尝试将一个任务也推送到通道中，若是通道容量满了，任务会堵塞，等待其他任务完成
	limiter <- struct{}{}

	fmt.Printf("worker id %d start\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("worker id %d done\n", id)
	// 任务已完成，释放一个令牌，允许其他任务继续执行
}

// ChannelSynchronization2 并发限流, 控制并发 goroutine 数量
func ChannelSynchronization2() {
	fmt.Println("demo: ChannelSynchronization2")

	// 创建一个容量为3的缓冲通道，用于控制并发 goroutine 数量，作为令牌桶
	limiter := make(chan struct{}, 3) // 限制并发 goroutine 数量为3

	wg.Add(10) // 总共有10个任务需要完成

	for i := 1; i <= 10; i++ {
		go worker(i, limiter)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("all workers done")
}
