/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-27 22:00:57
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-27 22:41:01
 * @FilePath: /src/chapter12/demo02/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	_ "sync"
)

func main() {
	c1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
			fmt.Println("aaaaa")
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c1)
	}

	fmt.Println("==============")

	c2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
			fmt.Println("bbbb")
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c2)
	}

	fmt.Println("==============")

	c3 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			c3 <- i
			fmt.Println("ccc")
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c3)
	}
}
