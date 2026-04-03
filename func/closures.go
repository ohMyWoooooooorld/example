package _func

import "fmt"

// 这是最经典的闭包用法，用于创建私有状态
func counter() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func Closures() {
	// 闭包
	// 闭包是一个函数值，它引用了其函数体之外的变量
	// 闭包可以访问和操作其函数体之外的变量
	// 闭包可以在其函数体之外被调用
	c1 := counter() // c1是一个闭包，它引用了num变量
	fmt.Println("c1() = ", c1())
	fmt.Println("c1() = ", c1())
	fmt.Println("c1() = ", c1())
	c2 := counter() // c2是一个闭包，它引用了num变量
	fmt.Println("c2() = ", c2())
	fmt.Println("c2() = ", c2())
	fmt.Println("c2() = ", c2())
}
