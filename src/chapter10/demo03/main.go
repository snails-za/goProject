/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2023-01-26 16:36:57
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2023-01-26 17:07:27
 * @FilePath: /src/chapter10/demo03/main.go
 * @Description: 面试题
 */
package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	// 这里实际上只创建了一个变量，所以stu地址是一样的，不能重复使用传递
	for _, stu := range stus {
		stu2 := stu
		m[stu.name] = &stu2
		fmt.Println(m)
	}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}
