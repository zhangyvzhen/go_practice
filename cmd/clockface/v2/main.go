package main

import (
	clockface "github.com/zhangyvzhen/go_practice/internal/math/v2"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
