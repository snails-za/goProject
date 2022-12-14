/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 17:22:37
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-13 23:42:19
 * @FilePath: /src/test/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 方式一：返回指针
	var nFlag = flag.Int("nflag", 4512, "help message for flag n")
	// 方式二：无返回值
	var flagvar int
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")

	flag.Parse()
	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Printf("%v\n", *nFlag)
	fmt.Printf("%v\n", flagvar)
}
