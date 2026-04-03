package _func

import "fmt"

func ChannelClose() {
	fmt.Println("demo: ChannelClose")

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("no more jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	close(jobs)
	fmt.Println("send all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("receive more jobs:", ok)
}
