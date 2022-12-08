/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 23:50:58
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-08 22:12:27
 * @FilePath: /src/chapter16/socket_tcp/client/main.go
 * @Description: TCP客户端
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [1024]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("收到server端发来的数据：", string(buf[:n]))
	}
}
