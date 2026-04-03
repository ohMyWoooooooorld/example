package _func

import "fmt"

// ChannelSelect1 select: 检测通道内是否有数据，有则接收，没有就执行其他任务
func ChannelSelect1() {
	fmt.Println("demo: ChannelSelect1")

	ch := make(chan string, 1)

	ch <- "Hello"

	// 通道内有数据，执行 case 分支
	select {
	case msg := <-ch:
		fmt.Println("成功接收通道数据:", msg)
	default:
		fmt.Println("default 被执行，通道内无数据")
	}

	// 通道内无数据，执行 default 分支
	select {
	case msg := <-ch:
		fmt.Println("成功接收通道数据:", msg)
	default:
		fmt.Println("default 被执行，通道内无数据")
	}
}
