/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-08 22:53:07
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-08 23:24:15
 * @FilePath: /src/chapter17/iolearn/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
