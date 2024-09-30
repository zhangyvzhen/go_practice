package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

// 现在需要你写一个程序，从 3 开始依次向下，当到 0 时打印 「GO!」 并退出，要求每次打印从新的一行开始且打印间隔一秒的停顿。
// 3
// 2
// 1
// Go!

//下面是我们如何划分工作和迭代的方法：
//
// 首先实现 打印 3
// 然后实现 打印 3 到 Go!
// 最后实现 在每行中间等待一秒
// 在测试的支持下，将功能切分成小的功能点，并使其首尾相连顺利的运行。

func TestCountdown(t *testing.T) {
	t.Run("Print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`
		//反引号语法是创建 string 的另一种方式，但是允许你放置东西例如放到新的一行

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}

	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
