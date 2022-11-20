/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 17:38:24
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 17:40:11
 * @FilePath: /src/chapter09/forrange/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main() {
	cities := make(map[string]string)
	cities["no1"] = "shanghai"
	cities["no2"] = "beijing"
	fmt.Println(cities)
	for k, v := range cities {
		fmt.Println(k, v)
	}
}