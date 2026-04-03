package _func

import (
	"fmt"
	"maps"
	"sort"
)

//go map 是一种键值对的集合
// map 是引用类型，它的默认值是 nil
// map 可以使用 make 函数来创建
// map 的键必须是可比较的类型，比如整数、浮点数、字符串、指针、通道、结构体等
// map 的值可以是任意类型

// map 的键值对是无序的，每次遍历的顺序都可能不同

// map 在循环中不可以进行删除和添加操作，添加和删除可能触发扩容/结构重组
// map 的并发问题，map并发不安全，可能会引起数据竞争问题
// 解决方法：
// 1. 使用 sync.Map 来替代 map，它是并发安全的
// 2. 在并发场景下，使用互斥锁（sync.Mutex）来保护对 map 的访问

// var m map[int]string 零值，不能直接使用，需要使用 make 函数来创建，添加元素会panic

func Map() {
	m := make(map[string]int, 3)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	for key, value := range m {
		// map是无序的
		fmt.Printf("key = %s, value = %d\n", key, value)
	}

	// 删除键值对
	delete(m, "b")
	fmt.Println("m = ", m)

	// 添加
	m["d"] = 4
	fmt.Println("m = ", m)

	fmt.Println("len(m) = ", len(m))

	clear(m) //用于清空map 和slice。 对于map，它会删除所有元素；对于slice，它会将所有元素设置为该类型对应的零值
	fmt.Println("map:", m)

	value, exist := m["a"] //Comma-ok 语法，用于检查 map 中是否存在指定的键
	if exist {
		fmt.Println("v = ", value)
	} else {
		fmt.Println("a is not in map")
	}

	//var m1 map[string]int
	//m1["error"] = 100 //运行时panic：assignment to entry in nil map，因为 m1 是 nil  map，没有分配内存空间

	fruit := make(map[string]int, 5)
	fruit["apple"] = 5
	fruit["orange"] = 10
	fruit["banana"] = 2
	fruit["grape"] = 8
	fruit["watermelon"] = 1

	// 遍历map，将键存储到切片中
	keys := make([]string, 0, len(fruit))
	for key, _ := range fruit {
		keys = append(keys, key)
	}
	fmt.Println("keys = ", keys)
	fmt.Println("fruits = ", fruit)

	// 对键进行排序
	sort.Strings(keys) //从小到大
	fmt.Println("sorted keys = ", keys)

	for _, key := range keys {
		// 按排序后的键输出对应的值
		fmt.Printf("key = %s, value = %d\n", key, fruit[key])
	}

	// 检查两个map是否相等
	if maps.Equal(fruit, map[string]int{"apple": 5, "orange": 10, "banana": 2, "grape": 8, "watermelon": 1}) {
		fmt.Println("fruit and the other map are equal")
	} else {
		fmt.Println("fruit and the other map are not equal")
	}
}
