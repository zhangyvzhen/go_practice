package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// 请注意，如果删除注释 「//Output: 3」，示例函数将不会执行。虽然函数会被编译，但是它不会执行。
func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// OutPut: 3
}
