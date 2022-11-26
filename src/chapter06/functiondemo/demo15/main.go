/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-26 16:24:15
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-26 16:26:01
 * @FilePath: /src/chapter06/functiondemo/demo15/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	demo()(55)
}
func demo()(func(int))  {
	return func (a int)  {
		fmt.Println(a)
	}
}