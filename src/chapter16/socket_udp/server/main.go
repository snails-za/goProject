/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-07 23:49:06
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-08 00:16:22
 * @FilePath: /src/chapter16/socket_tcp/server/main.go
 * @Description:UDP服务端
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
