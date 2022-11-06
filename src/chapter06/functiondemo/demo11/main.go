/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-06 20:20:17
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-06 20:29:13
 * @FilePath: /src/chapter06/functiondemo/demo11/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
var a = test()
func test() int {
	fmt.Println("全局变量定义！。。。")
	return 90
}
func main()  {
	fmt.Println("main()...")
}
func init()  {
	fmt.Println("init()...")
}