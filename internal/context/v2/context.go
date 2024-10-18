package context2

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求中获取 context.Context 对象，Context 是一个 Go 内置的机制，常用于处理请求的取消、超时等场景。
		// ctx 可以被用来检测客户端是否已经取消了请求。
		ctx := r.Context()

		//创建一个缓冲通道 data，用于在 Goroutine 和主函数之间传递数据。
		//缓冲大小为 1，意味着它可以存储一个值，而无需等待消费者接收
		data := make(chan string, 1)

		// 启动一个新的 Goroutine，它从 store.Fetch() 获取数据，并将结果发送到 data 通道中。
		// store.Fetch() 很可能是一个耗时的操作（例如从数据库或远程 API 获取数据）
		go func() {
			data <- store.Fetch()
		}()

		//select 语句用来监听多个通道（channel）的消息。
		//在这个例子中，代码监听两个情况：
		//一个是获取数据的 data 通道，
		//另一个是请求的 context.Context 的 Done() 通道（用于检测请求取消或超时）。
		select {
		// 如果 data 通道中有值（即数据成功从 store.Fetch() 获取并传递过来），处理函数将值写入 HTTP 响应中，通过 fmt.Fprint(w, d) 向客户端返回获取的数据。
		case d := <-data:
			fmt.Fprint(w, d)
		// 如果 ctx.Done() 通道关闭，表示请求被取消或超时。此时，触发 store.Cancel() 函数，取消或清理未完成的 store.Fetch() 操作。ctx.Done() 通常由客户端中止连接或者请求超时触发。
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
