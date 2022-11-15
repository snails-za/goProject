/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-16 00:06:39
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-16 00:38:50
 * @FilePath: /src/chapter06/exceptfunc/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"errors"
)
func test()  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	var num1 int = 10
	var num2 int = 0
	fmt.Println(num1 / num2)
}

func readConf(filename string) (err error) {
	if filename == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误！")
	}
}

func test2()  {
	err := readConf("config2.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test2继续执行。。。")
}

func main()  {
	test()
	fmt.Println("test()执行后，继续执行其他代码！")

	test2()
	fmt.Println("test2()执行后，继续执行其他代码！")
}