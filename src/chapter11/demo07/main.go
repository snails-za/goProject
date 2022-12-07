/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 01:49:03
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-07 01:55:44
 * @FilePath: /src/chapter11/demo06/main.go
 * @Description: 空接口的应用
 */
package main

import "fmt"

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	// 断言【这里不想解释，后边会有详细说明】
}
