package main

import (
	"github.com/zhangyvzhen/go_practice/internal/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second}
	mocking.Countdown(os.Stdout, sleeper)
}
