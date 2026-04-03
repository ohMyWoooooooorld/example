package _func

import (
	"fmt"
	"math"
)

const name = "张三"

func Constants() {
	fmt.Println(name)

	const n = 100

	const d = 3e20 / n
	fmt.Println(d)

	// 常量类型转换
	fmt.Println("d = ", int64(d))

	fmt.Println("sin(π/2) = ", math.Sin(math.Pi/2))

	//itoa 枚举器语法糖,与const配合使用，枚举器从0开始, 每个枚举器的值为上一个枚举器的值+1, 除非显式赋值,与行数有关
	const (
		a1 = iota
		a2 = 100
		a3 = iota
	)
	fmt.Printf("a1 = %d , a2 = %d , a3 = %d\n", a1, a2, a3)

	const (
		A, B = iota, iota
		C, D = iota, iota
	)
	fmt.Printf("A = %d , B = %d , C = %d , D = %d\n", A, B, C, D)

	const (
		Monday  = 77
		Tuesday = iota
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Printf("Monday = %d , Tuesday = %d , Wednesday = %d ,Thursday = %d , Friday = %d , Saturday = %d , Sunday = %d\n",
		Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	const (
		_  = iota
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Printf("c1 = %d , c2 = %d , c3 = %d", c1, c2, c3)
}
