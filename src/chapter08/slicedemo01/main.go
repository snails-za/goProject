/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 12:16:01
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 12:39:29
 * @FilePath: /src/chapter08/slicedemo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	// 1. 引用数组
	var arr = [5]int{1,2,3,4,5}
	var slice []int
	slice = arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 2. 通过make来创建
	var slice2 []int = make([]int, 5, 10)
	fmt.Println(slice2)

	// 3. 直接定义
	var slice3 []int = []int{1,2,3}
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}