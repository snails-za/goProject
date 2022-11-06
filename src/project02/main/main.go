/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-06 18:20:12
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-06 18:55:55
 * @FilePath: /src/project02/main/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"project02/model"
)

func main() {
	fmt.Println(model.Heroname)
	model.DbConnect()
}