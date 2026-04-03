package _func

import (
	"fmt"
	"time"
)

//核心作用：将时间转化为通道信号
//Go 中的 Timers 是一种机制，用于在未来的某个时间点或周期性地，向一个 time.Time 类型的通道发送一个值。
//这个值本身不重要，重要的是 信号的到达时间。

//1. 实现超时控制 (Timeout)
//这是 Timers 最常见的用途。通过将 time.After 或 time.NewTimer 的通道放入 select 中，可以确保一个操作不会永久阻塞。
//一旦计时器通道就绪，select 就会选择它，从而放弃当前正在等待的操作。
//2. 实现周期性任务 (Scheduling)
//使用 time.NewTicker 可以创建一个周期性的信号源，非常适合实现类似“每隔 5 分钟刷新一次缓存”、“每 1 小时记录一次系统状态”等任务。

/**
1. time.After 导致的 Goroutine 泄漏
这是最常见也最危险的错误。如果在超时前任务完成，time.After 产生的信号永远不会被读取，导致与之绑定的 Goroutine 悬挂。

正确方案： 如果任务需要被取消或提前完成，使用 time.NewTimer 并调用 Stop()。

2. 重置 Timer 时的通道残留（Race Condition）
当你调用 timer.Reset(d) 时，如果定时器已经过期，并且信号已经在通道 timer.C 中，但尚未被读取，Reset 会返回 true。此时，旧的信号仍然存在。

风险： 接收者可能会在下次 select 中立即收到这个 旧的、过期的信号，导致逻辑错误。

安全做法（在 Reset 前清空通道）：
// 检查定时器是否已经触发，若触发则清空通道
if !t.Stop() {
    select {
    case <-t.C: // 尝试取走通道中的残留信号
    default:
        // 通道已空
    }
}
t.Reset(newDuration)

3. Ticker 的时钟漂移（Jitter）
Ticker 信号的时间间隔并不总是 绝对精确 的。

原因： 信号的发送和接收都受 Go 调度器和当前系统负载的影响。

规则： Ticker 保证的是 最小间隔，但不保证准确的周期性。如果任务执行时间超过了 Ticker 间隔，信号可能会堆叠或延迟。
如果需要绝对精确的时钟，应该使用更底层的系统 API 或手动计算时间差。
*/

func Timers() {
	time1 := time.NewTimer(2 * time.Second)

	<-time1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

	resultCh := make(chan string, 1)

	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop() // 确保函数退出时停止计时器

	select {
	case <-resultCh:
		// 成功接收，此时必须停止定时器
		timer.Stop()
	case <-timer.C:
		fmt.Println("超时")
	}
}
