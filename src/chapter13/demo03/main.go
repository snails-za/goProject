/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-01 22:34:32
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-01 23:13:19
 * @FilePath: /src/chapter13/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}
type Student struct {
	Class string
	User
}

func (u User) SayName(name string) {
	fmt.Println("我的名字叫做", name)
}
func main() {
	u := User{
		Name: "wj",
		Age:  29,
		Sex:  true,
	}
	// check(u)
	s := Student{
		Class: "三年级二班",
		User:  u,
	}
	check(&s)
}
func check(inter interface{}) {
	v := reflect.ValueOf(inter)
	e := v.Elem()
	e.FieldByName("Class").SetString("四年级二班")
	fmt.Println(v)

}
