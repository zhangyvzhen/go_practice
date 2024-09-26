package arrays

import (
	"reflect"
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

	//t.Run("collection of any size", func(t *testing.T) {
	//	numbers := []int{1, 2, 3}
	//
	//	got := Sum(numbers)
	//	want := 6
	//
	//	if got != want {
	//		t.Errorf("got %d want %d given, %v", got, want, numbers)
	//	}
	//})

	//在本例中，针对该函数写两个测试其实是多余的，因为切片尺寸并不影响函数的运行
}

// 需要一个 SumAll 函数，它接受多个切片，并返回由每个切片元素的总和组成的新切片
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// 把每个切片的尾部元素相加（尾部的意思就是除去第一个元素以外的其他元素）
func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
		want := []int{5, 10}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 1})
		want := []int{0, 10}
		checkSums(t, got, want)
	})
}
