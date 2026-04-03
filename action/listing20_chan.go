package action

import (
	"fmt"
	"math/rand"
)

// 无缓冲通道模拟网球比赛
// 无缓冲通道若是没有接收者，将会阻塞，数据也不会进入发送通道

func player(name string, count chan int) {
	defer wg.Done()

	for {
		// 等待球被打过来, 若是球没有被发送到通道中，count会阻塞
		ball, ok := <-count
		// 使用 chan 一定要判断通道是否关闭，使用关闭的通道发送数据会导致 panic
		if !ok {
			// ok 判断通道是否关闭和是否有数据，若是关闭且没有数据，表示当前name赢了
			fmt.Printf("winner: %s\n", name)
			return
		}

		// 选随机数，然后用这个数来判断是否丢球
		n := rand.Intn(100)
		if n%17 == 0 {
			fmt.Printf("player %s miss the ball\n", name)

			// close 函数能关闭通道，通道被关闭后，还有数据残留在通道中仍然会被发送
			close(count)
			return
		}

		// 显示击球数，并将球加一
		fmt.Printf("player %s play the ball: %d\n", name, ball)
		ball++

		//将球打向对手
		count <- ball
	}
}

func GoroutineChan1() {
	// 创建一个无缓冲通道
	ball := make(chan int)

	wg.Add(2)
	go player("A", ball)
	go player("B", ball)

	ball <- 1

	wg.Wait()
}
