package _func

import "fmt"

// 泛型的提出解决两个问题
// 1. 代码因类型而重复编写
// 2. 解决类型安全来操作集合

// comparable 是 Go 泛型中引入的一个预定义的接口约束，它在 Go 语言中具有非常特定的含义和作用。
// comparable 约束定义了一组类型，这些类型在 Go 语言中是可以使用相等运算符 == 和不等运算符 != 进行比较的。
// 默认可比较的： 布尔型 (bool) 数字类型 (int, float64, complex128 等所有数字类型)
// 			    字符串 (string) 指针 (*T) 通道 (chan T)
//  			接口 (interface{}) 数组 ([N]T)：前提是数组的元素类型 T 也是可比较的。
// 				结构体 (struct)：前提是结构体的所有字段都是可比较的。
//以下类型不支持使用 == 或 != 直接比较，因此它们不满足 comparable 约束：
//切片（Slices）： 不能直接比较两个切片是否相等。
//映射（Maps）： 不能直接比较两个 map 是否相等（只能与 nil 比较）。map的键是可以比较的
//函数（Functions）： 函数值不能比较（只能与 nil 比较）。

func Contains[T comparable](slice []T, element T) bool {
	// 遍历切片，检查是否包含指定元素
	for _, item := range slice {
		// 只有 T 是 comparable 类型，才能使用 == 运算符
		if item == element {
			return true
		}
	}
	return false
}

// Stack 结构体是泛型的，它的元素类型由类型参数 T 决定
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.elements) - 1
	item := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return item, true
}

func Generics() {
	intSlice := []int{10, 20, 30}
	strSlice := []string{"apple", "banana", "cherry"}
	fmt.Printf("intSlice contains 20: %v\n", Contains(intSlice, 20))
	fmt.Printf("strSlice contains \"banana\": %v\n", Contains(strSlice, "banana"))

	//解决类型安全问题
	intStack := Stack[int]{}
	intStack.Push(100)

	v, ok := intStack.Pop()
	if ok {
		fmt.Printf("弹出值: %d (类型安全，无需断言)\n", v)
	}

	strStack := Stack[string]{}
	strStack.Push("go")

	s, ok := strStack.Pop()
	if ok {
		// 译器知道 s 已经是 string，不需要 s.(string) 断言
		fmt.Printf("弹出值: %s (类型安全，无需断言)\n", s)
	}
}
