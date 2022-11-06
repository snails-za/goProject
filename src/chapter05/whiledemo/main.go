/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-06 16:07:54
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-06 16:09:02
 * @FilePath: /go_code/chapter05/whiledemo/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}