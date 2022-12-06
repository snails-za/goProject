/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-03 23:38:01
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-06 23:53:54
 * @FilePath: /src/chapter15/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(a)

	var b map[string]string
	b = make(map[string]string)
	b["name"] = "xx"
	fmt.Println(b)
}
