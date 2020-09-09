package main

import "fmt"

func main() {
	// 三元运算符
	var n int
	var i int = 10
	var j int = 10

	// 传统的三元运算
	// n = i > j ? i : j
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println("n=", n)
}