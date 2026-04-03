package action

import (
	"fmt"
	"sort"
	"time"
)

// 结构体用于记录比赛结果
type RaceResult struct {
	TeamName   string
	TotalTime  time.Duration
	FinishTime time.Time
}

// RunV2 (改进版): 使用循环，一个 Goroutine 跑四棒
func RunV2(teamName string, baton chan int, results chan RaceResult) {
	// 确保这个 Goroutine 退出时，计数器减一
	defer wg.Done()

	startTime := time.Now()

	// 使用 for range 优雅地接收接力棒
	for runner := range baton {
		if runner == 4 {
			// 终点线
			fmt.Printf("Team %s Runner 4 arrived at the finish line\n", teamName)

			totalTime := time.Since(startTime)

			// 将结果发送到 Result 通道
			results <- RaceResult{
				TeamName:   teamName,
				TotalTime:  totalTime,
				FinishTime: time.Now(),
			}
			// 最后一棒结束，Goroutine 退出
			close(baton)
			return
		}

		// 模拟奔跑
		time.Sleep(50 * time.Millisecond) // 缩短时间，方便测试
		fmt.Printf("Team %s runner %d is running\n", teamName, runner)

		// 传递给下一个队员 (由于 baton 现在是缓冲通道，发送不会阻塞)
		newRunner := runner + 1
		fmt.Printf("Team %s runner %d is passing the baton to runner %d\n", teamName, runner, newRunner)

		// 无缓冲通道会阻塞baton发送者，直到baton接收者准备好接收，而baton自己也是一个同时接收者，会阻塞等待发送者发送数据
		baton <- newRunner
	}
}

func GoroutineChan2Fixed() {
	fmt.Println("--- 接力比赛开始 (修复并计时排名) ---")

	// 1. 修复死锁：将 baton 通道改为容量为 1 的缓冲通道, 必须用缓冲通道，否则会阻塞
	// 允许发送操作完成后 Goroutine 立即进入下一次循环接收
	batonA := make(chan int, 1)
	batonB := make(chan int, 1)
	batonC := make(chan int, 1)
	batonD := make(chan int, 1)

	// 2. 结果通道：容量为 4，用于收集所有队伍的结果
	results := make(chan RaceResult, 4)

	// 3. wg.Add(4) 刚好对应 4 个 RunV2 Goroutine
	wg.Add(4)

	go RunV2("A_Team", batonA, results)
	go RunV2("B_Team", batonB, results)
	go RunV2("C_Team", batonC, results)
	go RunV2("D_Team", batonD, results)

	// 4. 启动比赛
	batonA <- 1
	batonB <- 1
	batonC <- 1
	batonD <- 1

	wg.Wait() // 等待所有队伍 Goroutine 完成

	// 5. 关闭结果通道，表示结果收集结束
	close(results)

	// 6. 收集并处理结果
	allResults := make([]RaceResult, 0, 4)
	for res := range results {
		allResults = append(allResults, res)
	}

	// 7. 计算排名 (根据 FinishTime)
	sort.Slice(allResults, func(i, j int) bool {
		return allResults[i].FinishTime.Before(allResults[j].FinishTime)
	})

	fmt.Println("\n--- 比赛结果 ---")
	for i, res := range allResults {
		fmt.Printf("%d. %s 耗时: %v\n", i+1, res.TeamName, res.TotalTime)
	}
}
