/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 19:26:54
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 19:37:35
 * @FilePath: /src/chapter09/mapslice/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
func main()  {
	// monsters用于存放所有妖怪的信息：name和age
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil{
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "20"
	}
	fmt.Println(monsters)
	monsters = append(monsters, map[string]string{"name": "玉兔", "age": "300",})
	monsters = append(monsters, map[string]string{"name": "玉兔2", "age": "300",})
	monsters = append(monsters, map[string]string{"name": "玉兔3", "age": "300",})
	fmt.Println(monsters)
}