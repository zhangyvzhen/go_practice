package _select

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select 类似于 switch 语句，但它的每个 case 都是一个 channel 操作
	// 可以同时等待多个 channel 的发送或接收操作，并且在其中一个操作准备好时执行对应的 case
	// select 语句会阻塞，直到其中一个 case 准备就绪。如果没有 case 准备就绪，且存在 default，则会执行 default 中的代码
	// 如果多个 case 同时准备就绪，select 会随机选择其中一个执行
	// 可以结合 time.After 创建超时机制，处理长时间未返回的情况
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		//time.After 会在定义的时间过后发送一个信号给 channel 并返回一个 chan 类型
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
