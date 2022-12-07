/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 01:41:26
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-07 01:44:17
 * @FilePath: /src/chapter11/demo05/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

type Mover interface {
	move()
}

type dog struct {
	name string
}
type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}
func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}
