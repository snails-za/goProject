/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-29 23:03:01
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-29 23:08:07
 * @FilePath: /src/chapter12/demo03/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	var readc <-chan int = c1
	var writec chan<- int = c1
	writec <- 1
	fmt.Println(<-readc)
}
