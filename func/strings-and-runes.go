package _func

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

//对于大量字符串拼接，应使用 strings.Builder 或 bytes.Buffer 来提高性能。
// strings.Builder 是一个可变字符串，用于高效拼接字符串。
// bytes.Buffer 是一个可变字节切片，用于高效拼接字节。

func StringsAndRunes() {
	iterations := 10000

	fmt.Println("--- 1. 低效拼接 (String Concatenation) ---")
	start1 := time.Now()

	s := ""
	for i := 0; i < iterations; i++ {
		s += "a" // 每次循环都创建一个新的字符串
	}

	duration1 := time.Since(start1)
	fmt.Printf("拼接 %d 次，耗时: %v\n", iterations, duration1)

	fmt.Println("\n--- 2. 高效拼接 (strings.Builder) ---")
	start2 := time.Now()

	var builder strings.Builder
	for i := 0; i < iterations; i++ {
		builder.WriteString("a") // 内部使用可增长的 []byte，减少内存分配
	}
	_ = builder.String() // 最后只创建一次字符串

	duration2 := time.Since(start2)
	fmt.Printf("拼接 %d 次，耗时: %v\n", iterations, duration2)

	str := "你好呀！中国文化博大精深。"
	fmt.Println("str的字节数为:", len(str))
	for i := 0; i < len(str); i++ {
		//输出的是字节的十六进制表示
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()
	// RuneCountInString 统计字符串的字符数
	fmt.Println("str的字符数为:", utf8.RuneCountInString(str))

	// range遍历字符串，每个字符都是一个rune
	for index, value := range str {
		fmt.Printf("index = %d , value = %c\n", index, value)
	}

	fmt.Println()
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}
}
