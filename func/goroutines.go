package _func

import (
	"fmt"
	"time"
)

/**
goroutine 的特性
1. 并发执行：goroutine 是 Go 语言的并发执行单位，每个 goroutine 都在独立的栈上运行，共享相同的地址空间。
2. 轻量级：goroutine 的创建和销毁成本非常低，一个程序可以同时运行成千上万个 goroutine。
3. 调度：goroutine 由 Go 运行时系统负责调度，开发者无需关注底层的线程管理。
4. 通信：goroutine 之间主要通过通道（channel）进行通信，这是一种安全的、类型化的通信机制，避免了传统的锁和条件变量的使用。
5. 并行执行：在多核处理器上，多个 goroutine 可以并行执行，充分利用多核性能。

GPM 模型
1. G（Goroutine）：Go 并发执行的实体，轻量级“协程”。
2. P（Processor）：逻辑处理器（或上下文）。介于 G 和 M 之间的调度单元。P 持有一个 G 的运行队列，并负责将 G 传递给 M 执行。
                  GOMAXPROCS 决定 P 的数量。
3. M（Machine）：OS 线程（操作系统线程）。负责执行代码的实体，由 OS 内核调度。Go 调度器将 G 调度到 M 上执行。
4. 系统调用 (Syscall)：Goroutine 执行 I/O 操作时，会涉及系统调用。
                     当 G 阻塞在系统调用上时，Go 调度器会将 G 和 M 分离，并将 P 调度给其他 M，确保其他 Goroutine 不受阻塞影响。
5. GOMAXPROCS：Go 运行时系统的一个全局变量，用于设置最大并发线程数（逻辑处理器数）。
               控制 Go 程序可以同时使用的 OS 线程（M）的最大数量，默认等于 CPU 核数。
               限制并发度，充分利用 CPU 资源，并防止调度器创建过多的 OS 线程。

API 接口
1. 启动 Goroutine：go funcName(args) // 或 go func() { ... }()
2. 让出 CPU 时间片：runtime.Gosched() // 让出当前的 M 给其他 Goroutine 运行，但不能保证当前 Goroutine 不会被再次立即调度。
3. 获取当前 Goroutine 的 ID：runtime.NumGoroutine() // 返回当前运行的 goroutine 数量
4. 获取 P 的数量：runtime.GOMAXPROCS(0) // 获取当前 P 的数量。一般不需要修改，使用默认值（CPU 核数）即可。

可以使用 go run -race 命令来检测数据竞争问题。
*/

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func Goroutines() {
	f("direct")

	// 启动一个新的 goroutine 来执行 f 函数
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("oh!!!!!!")

	// 等待 goroutine 执行完成, 否则主程序会提前退出, 导致 goroutine 没有机会执行
	time.Sleep(time.Second)
	fmt.Println("done")
}
