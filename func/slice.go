package _func

import (
	"fmt"
	"slices"
)

//slice本质是一个结构体，包含了指向底层数组的指针、切片的长度和容量
// 切片的长度是指切片中元素的数量
// 切片的容量是指切片底层数组的长度，从切片的第一个元素开始计算

// 切片的append操作会返回一个新的切片，新的切片的长度会增加1，容量会根据情况增加
// 如果新的切片的容量大于等于原切片的容量的2倍，那么新的切片的容量会增加一倍
// 如果新的切片的容量小于原切片的容量的2倍，那么新的切片的容量会增加到原切片的容量的2倍

// slice的make操作可以创建一个指定长度和容量的切片 eg：
// s := make([]int, 0, 5)
// 第一个参数是切片的类型
// 第二个参数是切片的长度
// 第三个参数是切片的容量
// 如果省略第三个参数，那么容量和长度相同

// s := make([]int, 0, 5) 和 ss := []int{0, 0, 0, 0, 0} 两者在 len 和 cap 以及后续的 append 行为上有本质区别。
// s := make([]int, 0, 5) 创建了一个长度为0，容量为5的切片，s[0]会触发panic，因为s[0]没有初始化
// ss := []int{0, 0, 0, 0, 0} 创建了一个长度为5，容量为5的切片，ss[0]到ss[4]都被初始化了

// 切片被传递到函数内部后，若是不接收函数返回的新切片，那么函数操作后的切片并不影响原来的切片的长度和容量，只是修改了底层数组

//切片表达式：
// s[low:high] 左开右闭，表示从切片 s 中索引从 low 到 high-1 的元素
// s[low:high:max] 表示从切片 s 中索引从 low 到 high-1 的元素
// 包含索引 low 对应的元素，不包含索引 high 对应的元素
// 如果省略 high，那么会一直截取到切片的末尾 s[low:]
// 如果省略 low，那么会从切片的第一个元素开始截取 s[:high]
// 如果省略 high 和 low，那么会截取整个切片 s[:]
// max 表示新切片的容量，不能大于原切片的容量，不能小于 low 对应的索引
// 如果省略 max，那么新切片的容量会和原切片的容量相同

func Slice() {
	// 切片声明
	var s []int
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	arr := [6]int{1, 2, 3, 4, 5, 6}
	s1 := arr[1:3]
	fmt.Println("s1 = ", s1)
	fmt.Println("len(s1) = ", len(s1))
	fmt.Println("cap(s1) = ", cap(s1))
	s1[0] = 100
	fmt.Println("s1 = ", s1)
	fmt.Println("arr = ", arr)
	s1 = append(s1, 97)
	s1 = append(s1, 98)
	s1 = append(s1, 99)
	s1 = append(s1, 100)
	fmt.Println("s1 = ", s1)
	fmt.Println("len(s1) = ", len(s1))
	fmt.Println("cap(s1) = ", cap(s1))
	fmt.Println("arr = ", arr)
	fmt.Println("arr(s1) = ", len(arr))
	fmt.Println("arr(s1) = ", cap(arr))

	// ss1在函数sliceRise被扩容了，s生成了新的底层数组，不影响ss1，所以ss1的长度和容量都没有改变
	// ss2还没有调用sliceRise之前append操作扩容了，有自己新的底层数组，len=3，cap=4
	// ss2调用sliceRise后，s使用append操作没有扩容，为1，2，3，0，自增后为2，3，4，1
	// s
	ss1 := []int{1, 2}
	ss2 := ss1
	ss2 = append(ss2, 3)
	sliceRise(ss1)
	sliceRise(ss2)
	fmt.Println(ss1, ss2)

	// 切片扩展后不会影响到原切片
	orderLen := 5
	order := make([]int, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("order = ", order)
	fmt.Println("pollorder = ", pollorder)
	fmt.Println("lockorder = ", lockorder)
	fmt.Println("len(order) = ", len(order))
	fmt.Println("cap(order) = ", cap(order))
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))

	str := []string{"233", "666", "99", "456", "777"}
	c := make([]string, len(str))
	// 切片复制,将str复制到c中,str和c的元素是独立的,修改c不会影响str
	copy(c, str)

	// 切片比较
	if slices.Equal(str, c) {
		fmt.Println("str and c are equal")
	}

	c[0] = "hello"
	fmt.Println("str = ", str)
	fmt.Println("c = ", c)

	// 多维切片，每个元素都是一个切片,每个切片的长度都不同
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			fmt.Printf("twoD[%d][%d]\n", i, j)
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD = ", twoD)

}

func sliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
