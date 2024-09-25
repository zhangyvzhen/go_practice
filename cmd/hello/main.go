package main

// .(点导入)将导入的包中的所有标识符（变量、常量、类型和函数）直接引入当前命名空间，这样就可以直接使用它们，而无需加上包名。
// _(下划线导入)导入包但不使用该包中的任何标识符。通常用于初始化包中的副作用，例如执行 init() 函数。
// 什么都不用的正常导入，导入方式只需使用包名，并且在代码中通过包名引用该包中的标识符
import (
	"fmt"
	. "github.com/zhangyvzhen/go_practice/internal/hello"
)

func main() {
	fmt.Println(Hello("World"))
}
