package _func

import (
	"fmt"
	"iter"
	"slices"
)

// 25. range over Iterators 迭代器遍历

/**
理解这个迭代器：
	当执行 for item := range mySet 时，
    1.Go 会自动调用 IntoIter() 获取迭代器，
    2.创建一个 yield 函数传给迭代器
每次循环时：
	迭代器内部：
		if !yield("apple") { return }  // ① 把"apple"传给外部循环
									   // ② 如果外部循环执行了break
	外部循环：
		for item := range mySet {
			fmt.Println(item)  // 收到"apple"
			if item == "banana" { break }  // 这会触发yield返回false
		}
*/

// Set 是一个泛型集合类型，用于存储唯一的元素

type Set[T comparable] struct {
	// 存储元素的映射，键为元素值，值为空结构体 struct{} 占用 0 字节
	data map[T]struct{}
}

// NewSet 传入items，创建一个新的Set
func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{
		data: make(map[T]struct{}),
	}
	for _, item := range items {
		s.data[item] = struct{}{}
	}
	return s
}

// IntoIter 单值迭代器返回一个迭代器函数，用于遍历 Set 中的元素
func (s *Set[T]) IntoIter() func(yield func(T) bool) {
	// 返回一个迭代器函数
	return func(yield func(T) bool) {
		for item := range s.data {
			// 调用 yield 产出元素 item
			// 如果 yield 返回 false (即外部 for range 循环使用了 break)，则立即返回
			if !yield(item) {
				return
			}
		}
	}
}

// SingleValueIterator 单值迭代器示例
func SingleValueIterator() {
	fmt.Println("单值迭代器")
	mySet := NewSet("apple", "banana", "cherry", "data")

	fmt.Println("正常迭代")
	for fruit := range mySet.data {
		fmt.Printf("%s, ", fruit)
	}
	fmt.Println()

	// 外部循环使用 break 提前结束迭代
	fmt.Println("使用 break 提前结束迭代")
	for fruit := range mySet.data {
		fmt.Printf("%s, ", fruit)
		if fruit == "cherry" {
			break
		}
	}
	fmt.Println()
}

// 键值迭代器示例

// Pair 键值对结构体
type Pair[K, V any] struct {
	Key   K
	Value V
}

type SortedMap[K comparable, V any] struct {
	// 存储键值对的切片，键为 K 类型，值为 V 类型
	data []Pair[K, V]
}

func (s *SortedMap[K, V]) Add(key K, value V) {
	s.data = append(s.data, Pair[K, V]{Key: key, Value: value})
}

func (s *SortedMap[K, V]) IntoIter() func(yield func(K, V) bool) {
	// 返回一个迭代器函数
	return func(yield func(K, V) bool) {
		for _, pair := range s.data {
			// 调用 yield 产出元素 item
			// 如果 yield 返回 false (即外部 for range 循环使用了 break)，则立即返回
			if !yield(pair.Key, pair.Value) {
				return
			}
		}
	}
}

func KeyValueIterator() {
	fmt.Println("键值迭代器")
	//在当前正式版本的 Go 语言中（包括最新的 Go 1.25/1.26 版本），
	//Go 编译器尚未内置对自定义结构体（如 SortedMap）的 for range 循环支持。
	//尽管您已经根据 Go 社区的 迭代器提案 编写了完美的 IntoIter 方法，
	//但 Go 编译器目前只允许在内置类型（切片、数组、字符串、映射、通道）上使用 for range 循环。
	myMap := SortedMap[string, int]{}
	myMap.Add("Alice", 21)
	myMap.Add("Bob", 30)
	myMap.Add("Charlie", 25)

	fmt.Println("name - age")
	// 1. 获取迭代器函数
	iteratorFn := myMap.IntoIter()

	// 2. 构造一个 lambda/匿名函数作为 yield 回调
	// 3. 将这个回调传入 iteratorFn 进行消费
	iteratorFn(func(name string, age int) bool {
		// yield 回调函数体就是你的循环体
		fmt.Printf("name: %s, age: %d \n ", name, age)

		// 返回 true 表示继续迭代
		return true
	})

	// ------------------------------------------------------------------

	fmt.Println("\n------------------------------------")

	// 演示中断逻辑（break 效果）
	fmt.Println("演示中断（只打印前两个）")
	count := 0
	iteratorFn = myMap.IntoIter() // 每次新的迭代都需要一个新的迭代器函数

	iteratorFn(func(name string, age int) bool {
		if count >= 2 {
			return false // 返回 false 告诉迭代器：停止产出！
		}
		fmt.Printf("name: %s, age: %d \n ", name, age)
		count++
		return true // 继续迭代
	})
	fmt.Println()
}

// 链表迭代器示例

// 链表元素结构体
type element[T any] struct {
	// 指向下一个元素的指针
	next *element[T]
	// 元素值
	val T
}

// List 链表结构体
type List[T any] struct {
	// 头节点和尾节点指针
	head, tail *element[T]
}

// Push 向链表尾部添加一个元素
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		// 如果链表为空，创建一个新节点作为头节点和尾节点
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		// 否则，将新节点连接到尾节点的 next 指针，并更新尾节点指针
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All 返回一个迭代器函数，用于遍历链表中的所有元素
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// getFib 返回一个迭代器函数，用于生成 Fibonacci 数列
func getFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		// 无限循环，直到外部 yield 函数返回 false
		for {
			// 产出当前的 Fibonacci 数
			if !yield(a) {
				return
			}
			// 计算下一个 Fibonacci 数
			a, b = b, a+b
		}
	}
}

func FibonacciIterator() {
	fmt.Println("Fibonacci 迭代器")
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Printf("%d, ", e)
	}

	// Collect 函数可以将迭代器产出的所有元素收集到一个切片中
	all := slices.Collect(lst.All())
	fmt.Printf("\nall: %v\n", all)

	// 生成 Fibonacci 数列并打印，直到超过 100
	for n := range getFib() {
		if n > 100 {
			break
		}
		fmt.Println(n)
	}
}
