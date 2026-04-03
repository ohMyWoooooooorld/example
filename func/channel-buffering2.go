package _func

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 生产-消费者模型

func Producer(task chan int) {
	// 随机生产10到15个任务发送到通道中
	number := rand.Int63n(5) + 10

	for i := 1; i <= int(number); i++ {
		task <- i
		fmt.Println("producer:", i)
	}

	// 只有生产者才能关闭通道，发送信号告诉所有的消费者任务已经发送完毕
	close(task)
	fmt.Println("producer: all tasks sent")
}

func Consumer(workerID int, tasks chan int) {
	defer wg.Done()

	fmt.Println("worker", workerID, "start")
	// for range循环会安全自动接收未知数量的任务，直到通道被关闭且缓冲区为空
	for task := range tasks {
		time.Sleep(100 * time.Millisecond) // 模拟处理时间
		fmt.Printf("worker %d: get task %d and done\n", workerID, task)
	}
	fmt.Println("worker", workerID, "finish")
}

func ChannelBuffer2() {
	// 生产-消费者模型
	tasks := make(chan int, 5) // 缓冲通道，容量为5

	// 监控3个消费者
	wg.Add(3)
	for i := 0; i < 3; i++ {
		// 让消费者一个个去接取任务
		go Consumer(i, tasks)
	}

	// 启动生成者
	go Producer(tasks)

	// 等待所有消费者完成
	wg.Wait()
}
