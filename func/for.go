package _func

import "fmt"

func For() {
	// for 循环
	for i := 1; i < 4; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	//while 用法
	a := 2
	for a < 10 {
		fmt.Printf("%d ", a)
		a++
	}
	fmt.Println()

	// range 用法
	for b := range 10 {
		fmt.Printf("%d ", b)
	}
	fmt.Println()

	// 循环遍历数组
	for index, value := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("index = %d , value = %d\n", index, value)
	}
	fmt.Println()

	// 循环遍历字符串
	var str = "你好呀，go"
	for index, value := range str {
		fmt.Printf("index = %d , value = %c\n", index, value)
	}
	fmt.Println()

	// 循环遍历map,map是无序的，每次遍历都会不太一样
	fruits := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"orange": "orange",
	}
	for key, value := range fruits {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
	fmt.Println()

	// 遍历通道
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 关闭通道, 否则会导致死循环

	for value := range ch {
		fmt.Printf("chan value: %d\n", value)
	}
	fmt.Println()
}
