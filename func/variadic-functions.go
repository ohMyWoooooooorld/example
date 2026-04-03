package _func

import "fmt"

// 可变参数函数
// 可变参数函数可以接受任意数量的参数
// 可变参数函数的参数类型必须是相同的

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func VariadicFunctions() {
	// 可变参数函数
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))

	nums := []int{1, 2, 3, 4}
	// 可变参数函数可以接受切片作为参数
	fmt.Println(sum(nums...))
}
