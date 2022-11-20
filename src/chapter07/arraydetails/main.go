/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 11:45:59
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 12:00:43
 * @FilePath: /src/chapter07/arraydetails/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func test(arr [3]int)  {
	arr[0] = 77
}
func test1(arr *[3]int)  {
	(*arr)[0] = 77
}
func main()  {
	// 1. 数组的长度不能动态变化
	arr1 := [3]int {1,2,3}
	// arr1[3] = 4
	fmt.Println(arr1)

	// 2. 不指定数组大小，并不是一个数组，而是一个切片slice
	var arr2 []int
	fmt.Printf("%T", arr2)

	// 3. 数组中的元素可以是任意类型，但是不能混用
	// arr1[2] = 4.4
	// 4. Go的数组是值类型，在默认情况下是值传递,因此会进行值拷贝，数组见不会互相影响
	arr3 := [3]int {11,22,33}
	test(arr3)
	fmt.Println(arr3)
	test1(&arr3)
	fmt.Println(arr3)
}