package _func

import "fmt"

// 递归：是一个函数自己调用自己的过程
// 1. 基本条件：递归调用必须要有一个基本条件，当满足基本条件时，递归调用就会停止。
// 2. 递归步骤 (Recursive Step)：函数如何把问题变小，然后调用自己去解决这个更小的问题。

// 阶层：n! = n * (n-1)!
func face(n int) int {
	if n == 1 {
		return 1
	}
	return n * face(n-1)
}

func Recursion() {
	// 递归调用
	fmt.Println(face(7))

	// 斐波那契数列：f(n) = f(n-1) + f(n-2)
	var fib func(n int) int // 先声明fib函数

	// 定义fib函数，给fib赋值一个匿名函数
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) * fib(n-2)
	}

	// 递归调用
	fmt.Println(fib(7))
}
