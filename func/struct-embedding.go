package _func

import "fmt"

type base struct {
	ID  int
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base(ID=%d, num=%d)", b.ID, b.num)
}

type container struct {
	base
	str string
}

func StructEmbedding() {
	c := container{
		base: base{
			ID:  100,
			num: 200,
		},
		str: "hello,world!",
	}
	fmt.Println(c.describe())

	// 访问 base 结构体的字段
	fmt.Println("c.ID = ", c.ID)
	fmt.Println("c.num = ", c.num)

	type describer interface {
		describe() string
	}

	// d 是一个 describer 接口类型的变量
	var d describer = c
	fmt.Println(d.describe())
}
