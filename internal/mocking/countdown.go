package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

// SpySleeper
// 监视器（spies）是一种 mock，它可以记录依赖关系是怎样被使用的。
// 它们可以记录被传入来的参数，多少次等等。
// 在我们的例子中，我们跟踪记录了 Sleep() 被调用了多少次，这样我们就可以在测试中检查它。
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	Duration time.Duration
}

func (s *ConfigurableSleeper) Sleep() {
	time.Sleep(s.Duration)
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	// 当CountdownOperationsSpy作为依赖注入时
	// 每次执行Sleep 其实都只是向Calls数组中  追加一个 sleep
	s.Calls = append(s.Calls, sleep)
}

// 同时实现了这个接口
//
//	type Writer interface {
//		 Write(p []byte) (n int, err error)
//	}
//
// 使其可以作为一种输出方式，类似 os.Stdout
func (s *CountdownOperationsSpy) Write(_ []byte) (n int, err error) {
	// 当CountdownOperationsSpy作为依赖注入时
	// 每次执行 Write ( fmt.Fprintln 的实现里有调用 Write) 其实都只是向Calls数组中  追加一个 write
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(write io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, err := fmt.Fprintln(write, i)
		if err != nil {
			return
		}
	}
	sleeper.Sleep()
	_, err := fmt.Fprint(write, finalWord)
	if err != nil {
		return
	}
}
