/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 12:50:40
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 15:23:54
 * @FilePath: /src/chapter08/slicedemo02/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	slice := []int{1,2,3}
	fmt.Println(slice)
	fmt.Println(&slice[0])

	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)
	fmt.Println(&slice[0])

	slice = append(slice, slice...)
	fmt.Println(slice)
	fmt.Println(&slice[0])

	str := "hello"
	str1 := []rune(str)
	str1[0] = '被'
	str = string(str1)
	fmt.Println(str)
}