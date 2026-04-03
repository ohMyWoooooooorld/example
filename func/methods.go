package _func

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// measure 测量 geometry 接口的面积和周长
func measure(g geometry) {
	fmt.Println("g is:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perimeter:", g.perimeter())
}

// detectType 检测 geometry 接口的具体类型
func detectType(g geometry) {
	switch v := g.(type) {
	case rect:
		fmt.Println("rect with width:", v.width, "height:", v.height)
	case circle:
		fmt.Println("circle with radius:", v.radius)
	default:
		fmt.Println("unknown type")
	}
}

func Methods() {
	r := rect{
		width:  10,
		height: 5,
	}

	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rap := &r
	fmt.Println("rap's area:", rap.area())
	fmt.Println("rap's perimeter:", rap.perimeter())

	c := circle{
		radius: 5,
	}

	measure(r)
	measure(c)

	detectType(r)
	detectType(c)
}
