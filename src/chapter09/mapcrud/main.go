/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 17:30:06
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 17:37:49
 * @FilePath: /src/chapter09/mapcrud/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	cities := make(map[string]string)
	// 新增
	cities["no1"] = "beijing"
	cities["no2"] = "shjanghai"
	// 修改
	cities["no2"] = "shanghai"
	// 删除
	delete(cities, "no1")
	// key不存在，不会报错
	delete(cities, "no4")
	fmt.Println(cities)
	// 查询
	val, ok := cities["no1"]
	fmt.Println(val, ok)
}