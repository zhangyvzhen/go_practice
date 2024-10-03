package concurrency

import (
	"io"
	"net/http"
	"time"
)

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

func OnlineWebsiteChecker(url string) bool {
	// 设置超时
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	// 发送 GET 请求
	resp, err := client.Get(url)
	if err != nil {
		return false // 如果发生错误，返回 false
	}
	// defer 用于延迟函数的执行。
	// 被 defer 修饰的函数会在包含它的函数执行结束时被调用，通常用于执行清理操作，比如关闭文件、释放资源等。
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body) // 确保在函数退出时关闭响应体

	// 返回响应状态码是否为 200 OK
	return resp.StatusCode == http.StatusOK
}
