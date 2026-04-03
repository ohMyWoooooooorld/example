package _func

import "fmt"

func Array() {
	// 定长：一旦声明后，运行中无法改变
	// 值类型：数组是值类型，赋值和传递操作会复制整个数组，而不是引用。所以
	//        传递一个数组成本比较高，修改副本不会影响到原始的数组。同时slice是引用类型，
	//        传递一个slice成本比较低，修改副本会影响到原始的slice。
	// 内存连续性：数组元素在内存中是连续的，索引访问为O(1)时间复杂度
	// 零值初始化：数组的元素会被初始化为其类型的零值，[5]int 会初始化为 [0 0 0 0 0]
	//           [5]string 会初始化为 [" " " " " " " " " "]

	fmt.Println("1.数组的声明、初始化，定长特性")
	var a [5]int = [5]int{10, 20}
	fmt.Printf("arr a：%v, len: %d\n", a, len(a))

	b := [...]string{"abc", "go", "demo"} //自动推断数组长度
	fmt.Printf("arr b：%v, len: %d\n", b, len(b))

	c := [3]int{}
	//c = a cannot use a (type [5]int) as type [3]int in assignment，[5]int 和 [3]int 是不同的类型，不能直接赋值
	fmt.Printf("arr c：%v, len: %d\n", c, len(c))

	d := [5]int{1: 100, 4: 500}
	fmt.Printf("arr d：%v, len: %d\n", d, len(d))

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 99
	fmt.Printf("originalArray: %v\n", originalArray) //original 不会被修改
	fmt.Printf("copiedArray: %v\n", copiedArray)

	modifyArray(originalArray)                                         //传递给函数的是数组的副本，不会影响到原始的数组，但是拷贝了整个数组给函数
	fmt.Printf("originalArray after modifyArray: %v\n", originalArray) //original 不会被修改

	fmt.Printf("b[:2]: %v\n", b[:2])

	e := [...]int{100, 3: 400, 500}
	fmt.Println("idx:", e)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			fmt.Println("i:", i, "j:", j)
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

func modifyArray(arr [3]int) {
	arr[0] = 77
	fmt.Printf("arr in modifyArray: %v\n", arr)
}
