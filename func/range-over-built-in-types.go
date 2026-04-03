package _func

import "fmt"

func RangeOverBuiltInTypes() {
	// 遍历数组
	s := []int{1, 2, 3, 4, 5}
	for i := range s {
		fmt.Println("s[i] = ", s[i])
	}

	m := map[string]int{
		"apple":      1,
		"banana":     2,
		"origin":     3,
		"watermelon": 4,
	}
	// 遍历map
	for key, value := range m {
		fmt.Println("key = ", key, "value = ", value)
	}

	str := "你好，go！"
	for index, value := range str {
		fmt.Printf("index = %d , value = %c\n", index, value)
	}
}
