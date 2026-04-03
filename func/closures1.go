package _func

import (
	"fmt"
	"sync"
)

// 当你在定义一个函数的参数或变量时，如果该类型是一个函数类型，你只需要指明它的类型签名（Signature），而不需要给参数命名。
// 所以handler func(int) 也可以写成 handler func(data int)
func process(nums []int, handler func(int)) {
	for _, n := range nums {
		handler(n)
	}
}

func Closures1() {
	sum := 0
	process([]int{1, 2, 3}, func(x int) {
		sum += x
	})
	fmt.Println(sum)
	process([]int{4, 5, 6}, func(x int) {
		sum += x
	})
	fmt.Println(sum)

	//定义后立即执行，用完就不再调用了
	iAmFunc := func(message string, n int) int {
		fmt.Println(message)
		return n + 1
	}("我是一个闭包函数", 5)
	fmt.Printf("iAmFunc 的类型是：%T\n", iAmFunc)
	fmt.Println("测试信息1", iAmFunc)
	fmt.Println("测试信息2", iAmFunc)
	// ④ 接收闭包：c1 和 c2 是两个独立的闭包实例
	c1 := generator() //这条语句不执行闭包，只是把闭包如何执行（即函数定义和它所捕获的状态）告诉了 c1
	c2 := generator()

	fmt.Println("c1 第一次调用:", c1())
	fmt.Println("c1 第二次调用:", c1())

	fmt.Println("c2 第一次调用:", c2())
	fmt.Println("c1 第三次调用:", c1())

	number1 := []int{1, 2, 3}
	wg := sync.WaitGroup{}

	//一个 for 循环内创建匿名函数（通常是为了启动 Goroutine）时，匿名函数捕获的是循环变量的引用，而不是它在当前迭代中的值副本。
	for _, n := range number1 {
		wg.Add(1)
		go func(x int) { // 接收一个参数 val
			// defer 语句会在函数返回前调用，确保 wg.Done() 被调用
			defer wg.Done()
			fmt.Printf("正确! 处理数字: %d\n", x)
		}(n) // 传入当前迭代的 n 的值
	}

	for _, n := range number1 {
		wg.Add(1)
		// 匿名函数捕获了循环体外部的同一个 'n' 变量的引用
		go func() {
			defer wg.Done()
			// 当 Goroutine 运行时，循环往往已经结束，n 的值是最终值（3）
			fmt.Printf("错误! 处理数字: %d\n", n)
		}()
	}
	wg.Wait()
}

// 外部函数：负责创建并返回闭包
func generator() func() int {
	// ① 外部变量：在 generator 函数的作用域内定义
	count := 0

	// ② 匿名函数：被返回，它捕获了 count
	return func() int {
		// ③ 闭包体：访问并修改了外部变量 count
		count++
		return count
	} // 匿名函数结束
} // generator 函数结束
