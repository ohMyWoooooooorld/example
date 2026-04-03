package _func

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// readOp 读取操作, key 要读取的键, resp 读取操作完成后的响应通道
type readOp struct {
	key  int
	resp chan int
}

// key 要写入的键, val 要写入的值, resp 写入操作完成后的响应通道
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func StatefulGoroutines() {
	fmt.Println("demo: StatefulGoroutines")

	// 统计读取操作和写入操作的次数
	var readOps uint64
	var writeOps uint64

	// 创建读取操作通道和写入操作通道
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 启动一个 goroutine 来处理读取和写入操作
	go func() {
		var state = make(map[int]int)
		for {
			select {
			// 处理读取操作, 从状态 map 中获取对应的值并发送到响应通道
			case read := <-reads:
				// 处理读取操作, 从状态 map 中获取对应的值并发送到响应通道
				read.resp <- state[read.key]

			// 处理写入操作, 将值写入状态 map 并发送成功响应
			case write := <-writes:
				// 处理写入操作, 将值写入状态 map 并发送成功响应
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for range 100 {
		go func() {
			for {
				// 随机生成读取操作
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}

				// 发送读取操作到读取操作通道
				reads <- read
				<-read.resp

				// 统计读取操作次数
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for range 10 {
		go func() {
			for {
				// 随机生成写入操作
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}

				// 发送写入操作到写入操作通道
				writes <- write
				<-write.resp

				// 统计写入操作次数
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.AddUint64(&readOps, 0)
	writeOpsFinal := atomic.AddUint64(&writeOps, 0)

	fmt.Println("readOps:", readOpsFinal)
	fmt.Println("writeOps:", writeOpsFinal)
}
