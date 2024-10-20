package main

import (
	"github.com/zhangyvzhen/go_practice/internal/math/vFinal/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
