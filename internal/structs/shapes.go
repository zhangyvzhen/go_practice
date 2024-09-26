package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Area
// go 没有重载，通过 方法 实现类型效果
// 声明一个 结构体的方法
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
