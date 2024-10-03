package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteCheck(_ string) bool {
	// 故意放慢速度,明确等待 20 毫秒
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteCheck, urls)
	}
}
