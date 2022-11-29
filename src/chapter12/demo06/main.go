/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-29 23:59:49
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-30 00:03:20
 * @FilePath: /src/chapter12/demo06/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main() {
	ch := make(chan int)
	var writec chan<- int = ch
	var readc <-chan int = ch

	go SetData(writec)
	ReadData(readc)

}
func SetData(writec chan<- int) {
	for i := 0; i < 10; i++ {
		writec <- i
	}
}
func ReadData(readc <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-readc)
	}
}
