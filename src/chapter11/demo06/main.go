/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 01:49:03
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-07 01:52:13
 * @FilePath: /src/chapter11/demo06/main.go
 * @Description: 空接口的定义
 */
package main

import "fmt"

func main() {
	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	// 空接口类型的变量可以存储任意类型的变量。
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}
