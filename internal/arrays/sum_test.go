package arrays

import (
	"testing"
)

// 创建一个 Sum 函数，它使用 for 来循环获取数组中的元素并返回所有元素的总和
func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		// %v 格式化输出数组
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}
}
