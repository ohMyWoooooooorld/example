package _func

import "fmt"

func Map1() {
	//安全地删除map中的元素
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
		"data":   4,
	}

	// 要删除的键列表
	keysToDelete := make([]string, 0)

	for key, value := range m {
		if value%2 == 0 {
			// 将要删除的键存入切片中
			keysToDelete = append(keysToDelete, key)
		}
	}

	// 遍历要删除的键列表，安全删除map中的元素
	for _, key := range keysToDelete {
		delete(m, key)
	}
	fmt.Println("m = ", m)
}
