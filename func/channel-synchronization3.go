package _func

import (
	"fmt"
	"time"
)

// 超时控制与退出信号（select）

func workingGoroutine(id int, stop chan struct{}) {
	defer wg.Done()

	fmt.Println("worker goroutine", id, "开始工作")

	for {
		select {
		case <-stop:
			fmt.Println("worker goroutine", id, "收到停止信号，退出工作")
			return
		case <-time.After(time.Millisecond * 300):
			fmt.Println("worker goroutine", id, "继续工作并完成了一次周期性检测")
		}
	}
}

func ChannelSynchronization3() {
	fmt.Println("demo: ChannelSynchronization3")

	stopSignal := make(chan struct{})

	// 启动两个worker goroutine
	wg.Add(2)
	go workingGoroutine(1, stopSignal)
	go workingGoroutine(2, stopSignal)

	time.Sleep(2 * time.Second)

	// 发送停止信号
	fmt.Println("发送退出停止信号")
	close(stopSignal)

	// 使用 Select 实现超时等待，避免主协程永久堵塞
	select {
	case <-time.After(time.Second * 3):
		fmt.Println("超时等待3秒，仍未收到worker完成信号，强制退出")
	case <-func() chan struct{} {
		// 使用一个匿名 Goroutine 和通道来包装 wg.Wait()
		waitCh := make(chan struct{})
		go func() {
			wg.Wait()
			close(waitCh)
		}()
		return waitCh
	}():
		fmt.Println("所有worker goroutine已安全退出")
	}
}
