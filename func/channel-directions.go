package _func

import "fmt"

// 发送消息到 pings 通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 从 pings 通道接收消息并发送到 pongs 通道
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirections() {
	fmt.Println("demo: ChannelDirections")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// 发送消息到 pings 通道
	ping(pings, "passed message")

	// 从 pings 通道接收消息并发送到 pongs 通道
	pong(pings, pongs)

	// 从 pongs 通道接收消息并打印
	fmt.Println(<-pongs)

}
