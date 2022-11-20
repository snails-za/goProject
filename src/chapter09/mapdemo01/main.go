/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 16:56:29
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 17:12:50
 * @FilePath: /src/chapter09/mapdemo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	// map的声明和注意事项
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "bob"
	a["no2"] = "bob2"
	a["no3"] = "bob4"
	a["no2"] = "bob3"
	fmt.Println(a)
}