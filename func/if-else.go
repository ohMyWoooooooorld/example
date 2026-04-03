package _func

import "fmt"

func IfElse() {
	if 7%2 == 0 {
		fmt.Println("7 是偶数")
	} else {
		fmt.Println("7 是奇数")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is divisible by 2")
	}

	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num is a single digit")
	} else {
		fmt.Println("num is a multi-digit number")
	}
}
