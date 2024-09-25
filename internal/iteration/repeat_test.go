package iteration

//为一个重复字符 5 次的函数编写测试

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"
	if repeated != expected {
		t.Errorf("Got %s, expected %s", repeated, expected)
	}
}

// 基准测试benchmarks , 检测代码的执行效率
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 7)
	fmt.Println(repeated)
	// OutPut: aaaaaaa
}
