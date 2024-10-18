package context1

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint 是 Go 语言标准库 fmt 包中的一个函数，用于将格式化的输出写入到指定的 io.Writer 接口（如文件、网络连接或 HTTP 响应等）。
		// 它类似于 fmt.Print，但不同的是，fmt.Fprint 允许你指定一个输出目标，而不是直接输出到标准输出
		fmt.Fprint(w, store.Fetch())
	}
}

// http.Handler
// http.Handler 是一个接口，它定义了所有处理 HTTP 请求的处理器（Handler）需要实现的方法
// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

// ServeHTTP
// ServeHTTP 是处理器必须实现的方法。它会在有 HTTP 请求时被调用。
// ResponseWriter：用于构造和发送 HTTP 响应。
// *Request：包含了 HTTP 请求的详细信息。

// 每一个实现了 http.Handler 接口的类型都可以用作处理 HTTP 请求。常见的处理器是函数、结构体等。

// http.HandlerFunc
// Go 提供了一个内置的适配器 http.HandlerFunc (适配器类型)，它将普通函数转换为 http.Handler
// 它实际上是一个类型别名
// type HandlerFunc func(ResponseWriter, *Request)
// http.HandlerFunc 的核心在于它实现了 http.Handler 接口的 ServeHTTP 方法，因此可以将一个普通的函数作为 HTTP 处理器使用
// 它的定义大致如下:
// func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//    f(w, r)  // 直接调用该函数本身
// }
