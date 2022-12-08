/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-09 00:08:03
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-09 00:11:53
 * @FilePath: /src/chapter16/http_test/client/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//resp, _ := http.Get("http://www.baidu.com")
	//fmt.Println(resp)
	resp, _ := http.Get("http://127.0.0.1:8001/go")
	defer resp.Body.Close()
	// 200 OK
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
