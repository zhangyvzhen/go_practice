package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		// 启用1000 个goroutine 增加计数
		wantedCount := 1000
		counter := NewCounter()

		// 使用 sync.WaitGroup，这是同步并发进程的一种方便的方式
		// 一个 WaitGroup 等待一组 goroutine 的完成。
		// 主 goroutine 调用 Add 来设置等待的 goroutine 的数量。
		// 然后每一个 goroutines 运行并在完成时调用 Done。
		// 同时，可以使用 Wait 来阻塞，直到所有的 goroutines 完成。

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		// 上面一系列sync.WaitGroup的操作都是为了让断言在正确的时间启动（所有goroutines运行完成）
		wg.Wait()

		// 互斥锁在第一次使用后不能被复制
		// Counter(按值传递) 给 assertCounter 时，它将尝试创建一个互斥对象的副本
		// 应该传指针
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
