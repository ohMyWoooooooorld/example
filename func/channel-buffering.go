package _func

import "fmt"

func ChannelBuffering() {
	message := make(chan string, 2)

	message <- "Hello, World!"
	message <- "Hello, Go!"
	//message <- "6666" 通道容量为 2，所以可以发送 2 条数据，第 3 条数据会阻塞

	fmt.Println(<-message)
	fmt.Println(<-message)
}
