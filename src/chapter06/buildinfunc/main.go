/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-15 23:36:35
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-15 23:46:40
 * @FilePath: /src/chapter06/buildinfunc/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	// len函数
	fmt.Println(len("nihao"))

	// new函数
	num1 := 100
	fmt.Printf("num1的类型：%T, num1的值：%v, num1的地址：%v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2的类型：%T, num2的值：%v, num2的地址：%v\n", num2, num2, &num2)
	*num2 = 100
	fmt.Printf("num2指向的值是：%v\n", *num2)
}