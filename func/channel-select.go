package _func

import (
	"fmt"
	"time"
)

//select 语句的工作方式类似于 switch 语句，但它的 case 目标是 通道操作（发送或接收），而不是值。
//它的核心功能是在运行时，从多个通道操作中选择一个已就绪的操作执行
// 1.select 允许一个 Goroutine 同时监听多个通道。当多个通道都准备好操作时，select 会随机选择一个执行，保证公平性。
// 2.阻塞: 如果所有 case 中的通道操作都未就绪（例如，所有接收通道为空，所有发送通道已满），select 语句将 阻塞，直到其中一个通道操作就绪。
//	 就绪: 通道操作就绪的定义：
//		接收操作 (<-ch): 当通道中有数据可接收时，就绪。
//		发送操作 (ch <- data): 当通道有空间或有接收者准备好接收时，就绪。
//		default: 永远就绪。
// 3.如果有 多个 case 同时就绪，select 会 随机 选择其中一个执行。这保证了在并发竞争中，任何一个 Goroutine 都不会因为优先级而被饿死（Starvation）。
// 4.default 语句是可选的。如果出现 default，且所有其他 case 都未就绪，则 select 会立即执行 default 分支，而不会阻塞。这常用于实现 非阻塞 I/O 或轮询。
// 5.当一个通道被关闭后，对它的接收操作会 立即就绪(!!!!注意)
// 6.select 语句中的 case 如果使用 nil 通道，该 case 将 永远不会就绪。
//   用途: 这是一个强大的技巧，用于动态地启用或禁用 select 中的某个分支，而无需重写整个 select 块。
// 7.time.After(duration) 函数返回一个 只接收通道 (<-chan time.Time)。在指定的 duration 过去后，这个通道会收到一个值，从而使 case 就绪。

func ChannelSelect() {
	fmt.Println("demo: ChannelSelect")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from c2:", msg2)
		}
	}

}
