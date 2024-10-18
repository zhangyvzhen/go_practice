package main

import "fmt"

func Test() {
	str := "12345"
	for i := 0; i < len(str); i++ {

		fmt.Println(string(str[i]))
	}
}

func main() {
	Test()
}
