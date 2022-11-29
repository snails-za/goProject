/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-29 23:48:30
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-29 23:54:24
 * @FilePath: /src/chapter12/demo05/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	// select 执行没有顺序
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	case <-ch3:
		fmt.Println("ch3")
	default:
		fmt.Println("都不满足！")
	}

}
