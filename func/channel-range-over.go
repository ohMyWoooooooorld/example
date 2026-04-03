package _func

import "fmt"

func ChannelRangeOver() {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
