/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 23:50:58
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-08 00:21:07
 * @FilePath: /src/chapter16/socket_tcp/client/main.go
 * @Description: TCP黏包客户端
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
	}
}
