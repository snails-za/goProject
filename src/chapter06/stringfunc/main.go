/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-10 23:45:12
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-13 22:43:37
 * @FilePath: /src/chapter06/stringfunc/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
) 
func main()  {
	// 统计字符串长度，按照字节len(str)
	// golang的统一编码为utf8，ascii的字符占用一个字节，一个汉字占用3个字节
	str := "hello北京"
	fmt.Println("str len=", len(str))

	// 字符串遍历，同时处理中文问题 r := []rune(str)
	r := []rune(str)
	for i:=0; i < len(r); i++ {
		fmt.Printf("字符：%c\n", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("12a3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}

	// 整数转字符串
	str = strconv.Itoa(123)
	fmt.Println("str: ", str)

	// 字符串转字节
	bytes := []byte("hello bytes")
	fmt.Printf("bytes: %v\n", bytes)

	// byte 转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str: %v\n", str)

	// 10进制转2，8，16进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的2进制是：%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是：%v\n", str)

	// strings 包
	fmt.Println(strings.Contains("seafood", "foods"))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println("Go" == "go")
	fmt.Println(strings.HasPrefix("chicken", "chi"))
	fmt.Println(strings.HasSuffix("chicken.txt", ".txt"))
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println(strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
	fmt.Println(strings.TrimLeft(" n a lone gopher", " "))
	fmt.Println(strings.TrimRight(" a lone gopher", " "))
	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToUpper("Gopher"))
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

}