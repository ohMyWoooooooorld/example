package _func

import (
	"fmt"
	"time"
)

func ChannelSelect2() {
	fmt.Println("demo: ChannelSelect2")

	ch := make(chan string)

	go func() {
		time.Sleep(1500 * time.Millisecond)
		ch <- "Slow Operation Result"
	}()

	select {
	case result := <-ch:
		fmt.Println("超时之前成功接收通道数据:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("超时，未成功接收通道数据")
	}

	//避免 Goroutine 泄漏，接收通道数据
	// 为了避免 Goroutine 泄露，我们最好还是接收掉剩余的结果
	// 思考：为什么 1.5s 后发送的数据，主协程没有阻塞？
	// 答：select 已经执行完毕并跳出，主协程继续执行。

	// 实际生产中会使用 context.WithTimeout 更优雅地取消慢速操作。
	result := <-ch
	fmt.Println("最终接收通道数据:", result)

}
