package _func

import (
	"fmt"
	"reflect"
)

func Variables() {
	// 变量声明
	var a int
	var b float32
	var c string
	var d bool

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	// 变量赋值
	a = 10
	b = 3.14
	c = "hello,world!"
	d = true

	// 变量打印
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	e := 3e2
	fmt.Println(e)

	// 变量类型推导 %T
	f := 3.14
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("a is of type %T\n", a)

	type MyStruct struct {
		Name string
		Age  int
	}

	data := MyStruct{
		Name: "张三",
		Age:  18,
	}

	// 变量类型反射
	// reflect.TypeOf() 可以获取变量的类型信息
	x := reflect.TypeOf(data)
	fmt.Println(x.Name())
	fmt.Println(x.Kind())

	y := reflect.TypeOf(e)
	fmt.Println(y.Name())
	fmt.Println(y.Kind())
}
