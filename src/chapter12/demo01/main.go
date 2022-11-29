/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-26 23:03:08
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-27 21:50:05
 * @FilePath: /src/chapter11/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import (
	"fmt"
	"sync"
)
 func main()  {
	// goroutine
	// 在调用一个方法的时候，在前面加一个go就是goroutine，他会让方法异步执行
	var wg sync.WaitGroup
	wg.Add(1)
	go Run(&wg)
	wg.Wait()
 }
 func Run(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("running!")
 }