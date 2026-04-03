package _func

import (
	"fmt"
	"time"
)

func ChannelTimeout() {
	fmt.Println("demo: ChannelTimeout")

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
