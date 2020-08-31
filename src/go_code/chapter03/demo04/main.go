package main

import "fmt"
// 变量的注意事项
func main()  {
	
	// 该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i", i)
	// i = 1.2 //int, 原因是不能改变数据类型

	// 变量在同一个作用域（在一个函数或者代码块）内不能重名
	// var i int = 59
	// i := 99

}