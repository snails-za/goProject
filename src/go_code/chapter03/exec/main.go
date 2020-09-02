package main

import "fmt"

func main() {
	var num int = 9
	fmt.Println("num address=", &num)
	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Println("num=", num)
	fmt.Printf("ptr 的类型是：%T", ptr)
}