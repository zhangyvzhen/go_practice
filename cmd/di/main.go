package main

import (
	"github.com/zhangyvzhen/go_practice/internal/di"
	"net/http"
	"os"
)

func main() {
	// 打印到标准输出
	di.Greet(os.Stdout, "Elodie")

	// 作为响应打印
	// 注意端口前面有个 :
	err := http.ListenAndServe(":5000", http.HandlerFunc(di.MyGreeterHandler))
	if err != nil {
		return
	}

}
