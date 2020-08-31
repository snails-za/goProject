package main

import "fmt"
func main()  {
	
	// string的基本使用
	var address string = "北京长城 110 hello world"
	fmt.Println("address=", address)

	// 字符穿一旦赋值了，字符串就不能修改了，在Go语言中字符串是不可变的
	// var str string = 'hello' // invalid character literal (more than one character)go
	// str[0] = "a" // cannot assign to str[0]go
	// fmt.Println("str=", str)

	// 字符串的两种表示形式
	// 1> 双引号， 会识别转义字符
	str2 := "abc\nabc"
	fmt.Println(str2)
	// 2> 反引号， 以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	str3 := `* Go语言的字符串使用UTF-8编码标识Unicode文本，这样Golang统一使用UTF-8编码
	* 字符穿一旦赋值了，字符串就不能修改了，在Go语言中字符串是不可变的
	* 字符串的两种表示形式`
	fmt.Println(str3)

	var str = "hello " + "world"
	str += " hahaha"
	fmt.Println(str)
	// 当一个拼接的操作很长时，怎么办，可以分行写， 需要将'+'保留在上一行
	var str4 = "hello " + "world" + "hello " + 
	"world" + "hello " + "world" + "hello " + 
	"world" + "hello " + "world"
	fmt.Println(str4)
}