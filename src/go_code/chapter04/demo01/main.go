package main

import "fmt"

func main() {

	// 重点讲解 /  %
	// 说明，如果运算的数都是整数，运算的结果是小数，去掉小数部分，保留整数部分
	fmt.Println(10 / 4)
	
	var n1 float32 = 10 / 4
	fmt.Println(n1)
	
	// 说明，如果需要保留小数部分
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)
	
	// 演示%
	// 看一个公式：a % b = a - a / b * b
	fmt.Println("10%3=", 10 % 3)
	fmt.Println("-10%3=", -10 % 3)
	fmt.Println("10%-3=", 10 % -3)
	fmt.Println("-10%-3=", -10 % -3)

}