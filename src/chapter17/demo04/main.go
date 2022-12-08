/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-08 23:35:05
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-08 23:39:50
 * @FilePath: /src/chapter17/demo04/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dirs, err := os.ReadDir("../demo03")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dirs {
		fmt.Println(v.Name)
		fmt.Println(v.IsDir())
		fmt.Println(v.Type())
		fmt.Println(v.Info())
	}
}
