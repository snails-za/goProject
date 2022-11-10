/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-07 23:15:31
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-07 23:23:49
 * @FilePath: /src/chapter06/functiondemo/demo13/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func sum(num1, num2 int) int {
	// 当函数执行到defer时候，暂时不执行，会单独建立一块defer栈
	// 当函数执行完毕之后，再从defer栈中按照先入后出的顺序出栈并执行
	defer fmt.Println("ok1 num1 = ", num1)
	defer fmt.Println("ok2 num2 = ", num2)
	res := num1 + num2
	return res
}
func main()  {
	res := sum(10, 20)
	fmt.Println("res = ", res)
}