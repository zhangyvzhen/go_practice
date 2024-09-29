package di

import (
	"fmt"
	"io"
	"net/http"
)

// Greet
// 在测试时，将write设为bytes.Buffer 检查内容是否正确
// 在实际应用时，将write设为 os.Stdout 打印到标准输出 或其他输出
func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name)
	// 我们想要测试是否生成了正确的内容
	// 我们的函数不需要关心在哪里打印，以及如何打印，所以我们应该接收一个接口，而非一个具体的类型
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return
	}
	// fmt.Fprintf 和 fmt.Printf 一样，只不过 fmt.Fprintf 会接收一个 Writer 参数，用于把字符串传递过去，而 fmt.Printf 默认是标准输出。
}

func MyGreeterHandler(w http.ResponseWriter, _ *http.Request) {
	Greet(w, "world")
}
