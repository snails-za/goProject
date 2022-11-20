/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 11:34:07
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 11:39:45
 * @FilePath: /src/chapter07/arraydemo02/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	var arr1 []string = []string{"宋佳", "无用", "卢俊义"}
	fmt.Println(arr1)
	for index, value := range arr1{
		fmt.Println(index, value)
	}
}