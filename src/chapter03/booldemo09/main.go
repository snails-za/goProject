package main

import (
	"fmt"
	"unsafe"
)
func main()  {
	var b = false
	fmt.Println("b=", b)
	// 注意事项
	// 1.bool类型占用存储空间是1个字节
	fmt.Printf("b占用存储空间是：%d", unsafe.Sizeof(b))
	// 2.bool类型只能取值true或者false
}