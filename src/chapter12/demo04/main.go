/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-29 23:42:42
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-29 23:46:30
 * @FilePath: /src/chapter12/demo04/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	c1 <- 4
	c1 <- 5
	// 使用for-range必须使用close关闭管道
	close(c1)
	for v := range c1 {
		fmt.Println(v)
	}
}
