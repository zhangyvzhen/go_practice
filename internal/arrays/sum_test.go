package arrays

import (
	"testing"
)

// 创建一个 Sum 函数，它使用 for 来循环获取数组中的元素并返回所有元素的总和
func TestSum(t *testing.T) {
	// [5]int{1,2,3,4,5} 数组的大小也属于类型的一部分，[4]int 和 [5]int 是不同的类型
	// []int{1,2,3,4,5} 不指定大小，这是一个切片 slice
	// 数组 和 切片 也不是一个类型
	// 下面使用的全是切片，不是数组

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			// %v 格式化输出数组或切片
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
