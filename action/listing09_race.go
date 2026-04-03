package action

import (
	"fmt"
	"runtime"
)

var (
	counter int64
)

func incrementCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// 退出当前 goroutine 并允许其他 goroutine 执行
		runtime.Gosched()
		fmt.Printf("Goroutine %d: value = %d\n", id, value)
		value++
		counter = value
	}
}

// Goroutines9 演示了 goroutine 之间的共享变量和竞态条件
func Goroutines9() {
	wg.Add(2)
	go incrementCounter(1)
	go incrementCounter(2)
	wg.Wait()
	fmt.Println("Counter:", counter)
}
