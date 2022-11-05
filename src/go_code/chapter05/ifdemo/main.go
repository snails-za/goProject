/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-05 23:08:30
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-05 23:17:21
 * @FilePath: /go_code/chapter05/ifdemo/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

func main()  {
	// var age int = 19
	// if age > 18 {
	// 	fmt.Println("您成年了！")
	// } else {
	// 	fmt.Println("你还小孩！")
	// }

	if age := 18; age > 18 {
		fmt.Println("您成年了！")
	} else if age == 18 {
		fmt.Println("您刚好成年！")
	}	else {
		fmt.Println("你还小孩！")
	}
}