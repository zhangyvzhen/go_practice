package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// ch := make(chan int) // 创建一个用于传输 int 类型的 channel
	resultChannel := make(chan result)

	// 将想要加快速度的那部分代码并行化，同时确保不能并发的部分仍然是线性处理。
	// 这里不能并发的部分是 向map 写入数据  （fatal error: concurrent map writes）
	// 使用 channel 在多个进程间通信。

	for _, url := range urls {
		// 开启 goroutine 的唯一方法就是将 go 放在函数调用前面
		// 经常使用 匿名函数 启动goroutine
		// 在声明的同时执行 —— 通过匿名函数末尾的 () 实现的
		// 在声明匿名函数时所有可用的变量也可在函数体内使用
		// 确保 u 的值固定为循环迭代的 url 值，重新启动 goroutine。u 是 url 值的副本，因此无法更改
		go func(u string) {
			// 发送数据到 channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// 从 channel 接收数据
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
