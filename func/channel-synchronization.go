package _func

import (
	"fmt"
	"time"
)

// 主张 “不要通过共享内存来通信；而应通过通信来共享内存”。
func worker1(done chan bool) {
	fmt.Print("worker start")

	// 模拟耗时操作
	time.Sleep(100 * time.Millisecond)
	// 任务完成，发送信号
	done <- true
}

func ChannelSynchronization() {
	done := make(chan bool, 1)
	go worker1(done)
	// 等待worker完成
	<-done
	fmt.Println("worker finish")
}
