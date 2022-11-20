/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 19:47:37
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 19:54:07
 * @FilePath: /src/chapter09/mapsort/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sort"
)
func main()  {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Println(map1[v])
	}
}