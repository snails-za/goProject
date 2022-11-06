package main

import (
	"fmt"
)
// 演示golang中指针的类型
func main() {
	// 基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println("i的地址=", &i)

	// 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	// var ptr *int = &num
	// 1.ptr 是一个指针变量
	// 2.ptr的类型 *int
	// 3.ptr本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr=%v \n", ptr)
	fmt.Printf("ptr的地址是%v \n", &ptr)
	fmt.Printf("ptr的指向是%v \n", *ptr)
	fmt.Printf("ptr的指向是%v \n", *ptr)

}