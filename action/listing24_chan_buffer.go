package action

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(task chan string, workerID int) {
	defer wg.Done()

	for {
		task, ok := <-task
		if !ok {
			// 通道关闭，任务完成
			fmt.Printf("worker %d: task channel closed\n", workerID)
			return
		}
		fmt.Printf("worker %d: processing task %s\n", workerID, task)

		// 随机 sleep 一段时间，模拟处理任务的时间
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Println("worker", workerID, "finished task", task)
	}
}

func GoroutineChanBuffer() {
	task := make(chan string, 5)

	wg.Add(2)
	go worker(task, 1)
	go worker(task, 2)

	task <- "task 1"
	task <- "task 2"
	task <- "task 3"
	task <- "task 4"
	task <- "task 5"

	close(task)

	wg.Wait()
}
