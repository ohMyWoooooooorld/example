package action

import (
	"fmt"
	"runtime"
	"sync"
)

func Goroutines() {
	runtime.GOMAXPROCS(2)

	// wg 等待所有 goroutine 完成
	var wg sync.WaitGroup
	// add 2 表示有 2 个 goroutine 要等待完成
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		// Done 表示匿名函数退出时通知 wg 等待组，等待数量减 1
		defer wg.Done()

		// 打印小写字母 a-z 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		// Done 表示匿名函数退出时通知 wg 等待组，等待数量减 1
		defer wg.Done()

		// 打印大写字母 A-Z 3 次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Wait for Goroutines to finish")
	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("\nFinish Goroutines")
}
