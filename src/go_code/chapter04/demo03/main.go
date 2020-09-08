package main

import "fmt"

func main() {
	// 演示关系运算符
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	flag := n1 > n2
	fmt.Printf("flag 的type 是 %T flag:%s", flag, flag)
}