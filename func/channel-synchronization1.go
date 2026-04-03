package _func

import (
	"fmt"
	"time"
)

// ChannelSynchronization1 无缓冲通道实现 goroutine 同步, 严格同步
func ChannelSynchronization1() {
	fmt.Println("demo: ChannelSynchronization1")

	done := make(chan struct{})

	// task Goroutine B
	go func() {
		fmt.Println("task B Goroutine start")
		// 模拟耗时操作
		time.Sleep(100 * time.Millisecond)

		// B 任务完成，发送信号, 只有通道准备好接收信号，才会继续执行后续代码
		done <- struct{}{}
		fmt.Println("task B Goroutine finish")
	}()

	fmt.Println("main Goroutine wait for task B Goroutine finish")
	// 等待 B 任务完成
	<-done
	fmt.Println("main Goroutine continue after task B Goroutine finish")
}
