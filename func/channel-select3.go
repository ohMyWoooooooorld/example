package _func

import (
	"fmt"
	"time"
)

func doWork(id int, dataCh <-chan int, stopCh <-chan struct{}) {
	defer wg.Done()

	fmt.Printf("worker %d: start\n", id)
	for {
		select {
		case data, ok := <-dataCh:
			if !ok {
				// dataCh 通道被关闭，任务处理完成了，worker 安全退出
				fmt.Printf("worker %d: dataCh closed, exit\n", id)
				return
			}
			fmt.Printf("worker %d: received data %d\n", id, data)
			time.Sleep(50 * time.Millisecond)
		case <-stopCh:
			// 接收到 stopCh 信号，worker 开始退出
			fmt.Printf("worker %d: stop signal received, exit\n", id)
			return
		}
	}
}

func ChannelSelect3() {
	fmt.Println("demo: demoGracefulShutdown")

	dataCh := make(chan int, 5)
	stopCh := make(chan struct{})

	// 发送停止信号，关闭 stopCh 通道，通知 worker 任务处理完成
	wg.Add(1)
	go doWork(1, dataCh, stopCh)
	for i := 1; i <= 10; i++ {
		dataCh <- i
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("send stop signal")
	stopCh <- struct{}{}

	// close dataCh 通道，通知 worker 任务处理完成
	wg.Add(1)
	go doWork(2, dataCh, stopCh)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("close dataCh\n")
	close(dataCh)

	wg.Wait()
	fmt.Println("all workers exit")

}
