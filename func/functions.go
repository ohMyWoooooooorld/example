package _func

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func Functions() {
	response := plus(1, 2)
	fmt.Println("response = ", response)

	res := plusPlus(1, 2, 3)
	fmt.Println("res = ", res)
}
