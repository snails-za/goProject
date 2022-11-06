package main

import (
	"fmt"
	"strconv"
)
// 演示go中string转成基本数据类型
func main(){
	str := "true" 
	var b bool
	// b, _ = strconv.ParseBool(str)
	// 说明
	// 1.strconv.ParseBool(str) 函数会返回两个值(value bool, err error)
	// 因为只想获取到value bool, 不想获取err 所以使用 _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b = %v \n", b, b)

	var str2 string = "123456789"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 := int(n1)
	fmt.Printf("n1 type %T n1=%v \n", n1, n1)
	fmt.Printf("n2 type %T n2=%v \n", n2, n2)
	
	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	f2 := float32(f1)
	fmt.Printf("f1 type %T f1=%v \n", f1, f1)
	fmt.Printf("f2 type %T f2=%v \n", f2, f2)

	// 注意：
	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T n3=%v", n3, n3)
}