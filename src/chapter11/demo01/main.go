/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-26 23:03:08
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-27 21:34:26
 * @FilePath: /src/chapter11/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
type Animal interface {
	Eat()
	Run()
}
type Cat struct {
	Name string
	Sex bool
}
type Dog struct {
	Name string
}
func (c Cat) Run() {
	fmt.Println(c.Name, "开始跑")
}
func (c Cat) Eat()  {
	fmt.Println(c.Name, "开始吃")
}

func (d Dog) Eat()  {
	fmt.Println(d.Name, "开始吃")
}
func (d Dog) Run()  {
	fmt.Println(d.Name, "开始跑")
}

var L Animal
func Myfunc(a Animal)  {
	// a.Run()	
	// a.Eat()	
	L = a
}
func main()  {
	var a Animal
	a = Cat{
		"cat",
		true,
	}
	a.Run()
	a.Eat()

	Myfunc(a)
	L.Run()
}