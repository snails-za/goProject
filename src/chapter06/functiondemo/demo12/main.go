/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-06 23:54:18
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-07 00:01:30
 * @FilePath: /src/chapter06/functiondemo/demo12/main.go
 * @Description: 匿名函数使用
 */
package main

import (
	"fmt"
)
func main()  {
	res := func (n1, n2 int) int {
		return n1 + n2
	}(10, 11)
	fmt.Println(res)

	f := func (a, b int) int {
		return a - b
	}
	fmt.Println(f(10, 5))
}