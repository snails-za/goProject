/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-08 23:09:24
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-08 23:23:49
 * @FilePath: /src/chapter17/demo02/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("hello.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	file.Write([]byte("你好"))
}
