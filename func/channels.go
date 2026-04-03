package _func

import "fmt"

// Channels 无缓冲通道
func Channels() {
	message := make(chan string)

	go func() {
		message <- "Hello, World!"
	}()

	// 有接收者，才能从通道中接收数据
	msg := <-message
	fmt.Println(msg)
}
