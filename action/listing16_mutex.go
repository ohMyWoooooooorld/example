package action

import (
	"fmt"
	"runtime"
	"sync"
)

// 互斥锁（Mutex）
// 互斥锁 用于在代码上创建一个临界区
// 互斥锁 能保证同一时间只有一个goroutine 可以执行临界区的代码
var mutex sync.Mutex

func incMutexCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻，只有一个 goroutine 能够获取到锁,进入临界区
		mutex.Lock()

		value := counter
		runtime.Gosched()
		value++
		counter = value

		// 释放锁，允许其他 goroutine 进入临界区
		mutex.Unlock()
	}
}

func GoroutineMutex() {
	wg.Add(2)
	go incMutexCounter(1)
	go incMutexCounter(2)
	wg.Wait()
	fmt.Println("counter:", counter)
}
