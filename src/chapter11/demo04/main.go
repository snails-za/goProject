/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 01:35:57
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-07 01:40:22
 * @FilePath: /src/chapter11/demo04/main.go
 * @Description: 类型与接口的关系
 */
package main

import "fmt"

type Sayer interface {
	say()
}
type Mover interface {
	move()
}
type dog struct {
	name string
}

func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	// 一个类型实现多个接口
	var x Sayer
	var y Mover
	d := dog{"旺财"}
	x = d
	y = d
	x.say()
	y.move()
}
