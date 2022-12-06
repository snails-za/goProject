/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 01:23:01
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-07 01:29:55
 * @FilePath: /src/chapter11/demo02/main.go
 * @Description: 指针接收者实现接口
 */
package main

import "fmt"

type Mover interface {
	move()
}

type dog struct{}

func (d *dog) move() {
	fmt.Println("狗会动")
}
func main() {
	var x Mover
	// var wangcai = dog{} // 旺财是dog类型
	// x = wangcai         // x不可以接收dog类型
	var fugui = &dog{} // 因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui,fugui是*dog类型。
	x = fugui          // x可以接收*dog类型
	x.move()
}
