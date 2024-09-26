package interfaces

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	// 表格驱动测试
	t.Run("Table driven tests", func(t *testing.T) {
		// 匿名的结构体
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			{"Rectangle", Rectangle{10.0, 10.0}, 100.0},
			{"Circle", Circle{10.0}, 314.1592653589793},
			{"Triangle", Triangle{12, 6}, 36.0},
			// 可以手动命名
			{name: "Triangle1", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		}

		for _, tt := range areaTests {
			// 写在t.Run里 并添加命名，方便定位
			t.Run(tt.name, func(t *testing.T) {
				got := tt.shape.Area()
				if got != tt.hasArea {
					t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
				}
			})
		}
	})
}
