package action

import (
	"fmt"
	"time"
)

// 无缓冲通道模拟4人接力比赛

func Run(teamName string, baton chan int) {
	runner, ok := <-baton
	if !ok {
		fmt.Printf("%s runner %d finished the race\n", teamName, runner)
		return
	}

	if runner == 4 {
		fmt.Printf("%s runner %d finished the race\n", teamName, runner)
		close(baton)
		wg.Done()
		return
	} else {
		// 递归启动进程，致命问题： 每跑一棒都会启动一个新的 Goroutine，但旧的 Goroutine 并未退出。
		// Goroutine 泄漏/堆栈爆炸
		go Run(teamName, baton)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("team %s runner %d is running\n", teamName, runner)

	newRunner := runner + 1
	fmt.Printf("team %s runner %d is passing the baton to runner %d\n", teamName, runner, newRunner)
	baton <- newRunner
}

func GoroutineChan2() {
	baton1 := make(chan int)
	baton2 := make(chan int)
	baton3 := make(chan int)
	baton4 := make(chan int)

	wg.Add(4)
	go Run("A", baton1)
	go Run("B", baton2)
	go Run("C", baton3)
	go Run("D", baton4)

	baton1 <- 1
	baton2 <- 1
	baton3 <- 1
	baton4 <- 1

	wg.Wait()
}
