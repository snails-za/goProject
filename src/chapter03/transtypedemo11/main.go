package main

import "fmt"
func main()  {
	var i int32 = 100
	// 希望将 i --> float
	var n1 float32 = float32(i)
	fmt.Println(n1)
	var n2 int8 = int8(i)
	fmt.Println(n2)
	var n3 int64 = int64(i) // 低精度 -> 高精度
	fmt.Println(n3)

	// 在转换中，比如将 int8 【-128 --- 127】， 编译时不会报错
	// 只是转换的结果是按溢出处理， 和我们希望的结果不一样
	var num1 int64 = 999999
	var num2 int8 = int8(num1) + 127
	fmt.Println(num2)
}