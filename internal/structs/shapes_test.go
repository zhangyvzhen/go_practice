package structs

import (
	"math"
	"testing"
)

// 编程计算一个给定高和宽的长方形的周长
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// 编写一个函数 Area(width, height float64) 来返回长方形的面积
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	//为圆形写一个类似的函数
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := math.Pi * 100
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
