package _func

import (
	"fmt"
)

// Pointer 有以下特性
// 1.类型安全
// 2.不支持指针运算
// 3.创建指针后，无需手动释放，gc会回收
// 4.go都是值传递，而不是引用传递，将地址传递后，可以通过指针修改原变量的值
// 5.指针未被初始化后，默认值为nil

// Pointer 有以下需要掌握的知识点
// 1. 两个核心操作符：&取地址符（用于获取一个变量的内存地址），*指针解引用符（用于获取指针所指向的变量的值。也可以通过它来修改指针指向的值）
// 2. 指针的声明
// 3. new函数，用于创建一个指针，指向一个新分配的内存空间
// 4. 指针与函数：在函数间共享和修改数据
// 5. 指针与结构体：可以直接使用 . 来访问结构体字段，而无需显式解引用

// 注意事项
// 1. 指针的零值为nil，不能对nil指针进行解引用操作，否则会panic，避免指针的空指针异常：在使用指针前，先检查是否为nil
// 2. 指针适用范围：函数传参、传递的数据结构很大用指针避免复制、表达一个变量为不可用时的nil指针
// 3. 注意指针的生命周期和作用域
// 4. 区分new和make：
//    new用于为各种类型分配内存，返回一个指向该类型零值的指针 *T；
//    make用于创建引用类型（如slice、map、channel），并初始化它们的内部结构，返回的是类型本身 T（而不是指针）

//a 是一个 int 类型的变量（一个信箱）。
//&a 是获取 a 的内存地址。这个地址值的 类型 是什么？因为它指向一个 int，所以它的类型是 *int（一个指向 int 的地址）。
//func modifyValueByPointer(ptr *int) 这个函数签名声明：“我需要一个参数，名字叫 ptr，它的类型必须是 *int”。
//因此，当你调用 modifyValueByPointer(&a) 时，你传递的 &a（一个 *int 类型的值）正好与函数期望的参数类型 *int 完美匹配。

type Phone struct {
	Brand string
	CPU   string
	RAM   int
}

type UserRequest struct {
	name string
	age  *int //使用指针来表示为可选字段
}

func Pointer() {
	x := 10
	fmt.Printf("x = %d, &x = %p\n", x, &x)

	var p *int
	p = &x
	fmt.Printf("p的地址是%p\n", p)
	fmt.Printf("p指向的值是%d\n", *p)

	*p = 200
	fmt.Printf("x = %d\n", x)
	fmt.Printf("new p = %d\n", *p)

	doubleValue(x)
	fmt.Printf("x = %d\n", x)
	doubleValueByPointer(&x)
	fmt.Printf("x = %d\n", x)

	myPhone := Phone{
		Brand: "Apple",
		CPU:   "A19",
		RAM:   8,
	}
	fmt.Printf("myPhone = %v\n", myPhone)
	upgradePhone(&myPhone)
	fmt.Printf("myPhone = %v\n", myPhone)

	var p1 *int
	fmt.Printf("p1 is %p\n", p1)
	// 指针的零值为nil，不能对nil指针进行解引用操作，否则会panic，避免指针的空指针异常：在使用指针前，先检查是否为nil
	if p1 == nil {
		fmt.Println("p1 is nil, it points to nothing")
	}

	p2 := new(int)
	fmt.Printf("p2 is %p, *p2 is %d\n", p2, *p2)

	*p2 = 100
	fmt.Printf("*p2 is %d\n", *p2)

	//空元素也是切片的一个元素，只是它的值为空字符串，
	myLang := make([]string, 0, 4)
	myLang = append(myLang, "Go", "Python")
	fmt.Printf("myLang = %v\n", myLang)
	addNewLang(myLang)
	fmt.Printf("myLang = %v\n", myLang)

	//使用切片等引用结构，最好使用指针传递
	myLang1 := make([]string, 0, 4)
	myLang1 = append(myLang1, "Go", "Python")
	fmt.Printf("myLang1 = %v\n", myLang1)
	addNewLangByPointer(&myLang1)
	fmt.Printf("myLang1 = %v\n", myLang1)

	//使用切片等引用结构，最好使用返回切片
	myLang2 := make([]string, 0, 4)
	myLang2 = append(myLang2, "Go", "Python")
	fmt.Printf("myLang2 = %v\n", myLang1)
	myLang2 = addNewLangByPointer2(myLang2)
	fmt.Printf("myLang2 = %v\n", myLang2)

	//nil 指针可以清晰地表达“这个值不存在”的语义
	req1 := UserRequest{
		name: "Alice",
	}
	processUserRequest(&req1)

	age := 0
	req2 := UserRequest{
		name: "",
		age:  &age,
	}

	processUserRequest(&req2)

}

func doubleValue(v int) {
	v = 2 * v
}

func doubleValueByPointer(ptr *int) {
	*ptr = 2 * *ptr
}

func addNewLang(myLang []string) {
	myLang = append(myLang, "Rust")
	myLang[1] = "JavaScript"
}

func addNewLangByPointer(myLang *[]string) {
	*myLang = append(*myLang, "Rust")
	(*myLang)[1] = "JavaScript"
}

func addNewLangByPointer2(myLang []string) []string {
	myLang = append(myLang, "Rust")
	myLang[1] = "JavaScript"
	return myLang
}

func upgradePhone(phone *Phone) {
	phone.RAM = 16
	fmt.Println("upgradePhone: RAM upgraded to 16GB")
}

func processUserRequest(req *UserRequest) {
	if req.age != nil {
		fmt.Printf("User %s is %d years old\n", req.name, *req.age)
	} else {
		fmt.Printf("User %s age is not provided\n", req.name)
	}
}
