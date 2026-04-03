package _func

import "fmt"

func Switch() {
	score := 95

	fmt.Println("without fall-through")
	switch score {
	case 100:
		fmt.Println("满分")
	case 95:
		fmt.Println("优秀")
		//自动break，不会执行下一个case
	case 90:
		fmt.Println("良好")
	default:

		fmt.Println("继续努力")
	}

	// 带fall-through的switch语句
	fmt.Println("with fall-through")
	switch score {
	case 100:
		fmt.Println("满分")
	case 95:
		fmt.Println("分数是95分")
		fallthrough //强制执行下一个case代码
	case 90:
		fmt.Println("分数大于90分")
		//没有fallthrough，不会执行下一个case，所以停止了
	default:
		fmt.Println("分数小于90分")
	}

	day := 3
	switch day {
	case 6, 7:
		fmt.Println("是周末")
	default:
		fmt.Println("是工作日")
	}

	//类型switch
	//检查一个接口类型变量所持有的实际具体类型
	printType(100)
	printType("hello")
	printType(true)
	printType(3.14)
	printType(3e20)
}

// 接收一个接口类型的参数，根据实际类型打印不同的信息
func printType(i interface{}) {
	// 类型断言,i.(type) 可以获取接口变量i的实际类型
	// 类型断言可以在switch语句中使用，用于判断接口变量的实际类型, 变量v的类型会被自动断言(type assertion)
	switch v := i.(type) {
	case int:
		fmt.Println("i 是 int 类型, 值为:", v)
	case string:
		fmt.Println("i 是 string 类型, 值为:", v, ", 长度为:", len(v))
	case bool:
		fmt.Println("i 是 bool 类型, 值为:", v)
	default:
		fmt.Printf("i 是未知类型, %T\n", v)
	}
}
