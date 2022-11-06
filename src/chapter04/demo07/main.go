package main

import "fmt"

func main() {
	// 要求：可以从控制台接收用户信息 
	// 方式1  fmt.Scanln
	// 先声明需要的变量
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名 ")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)
	fmt.Println("请输入薪资 ")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是 %v \n年龄是 %v \n薪水是 %v \n是否通过考试 %v \n", name, age, sal, isPass)
	
	// 方式2  fmt.Scanf
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v \n年龄是 %v \n薪水是 %v \n是否通过考试 %v \n", name, age, sal, isPass)
}