package sync

import "sync"

type Counter struct {
	// Mutex 是互斥锁。互斥锁的零值是一个解锁的互斥锁。
	mu sync.Mutex

	//// 另一种写法, 互斥锁嵌入到struct中, 这种写法“看起来”更优雅，但是是错误的。
	//// 嵌入类型 意味着该类型的方法成为 公共接口的一部分 。我们应该非常小心地使用我们的公共 API
	//// 当我们将某个对象公开时，其他代码也可以将自己与之结合。我们总是希望避免不必要的耦合。
	//// 如果你的类型的调用者开始调用这些方法，那么暴露 Lock 和 Unlock 往好里说就是混淆，往坏里说就是对你的软件有潜在的危害
	//sync.Mutex
	value int
}

// NewCounter
// Counter 中有互斥锁，在第一次使用后就不能复制,如果在函数中传递Counter默认是按值传递，会创建副本
// 使用Counter{} 在传指针时还要取地址，不如直接都用指针
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// 串行化这个方法
	c.mu.Lock()
	defer c.mu.Unlock()

	//c.Lock()
	//defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
