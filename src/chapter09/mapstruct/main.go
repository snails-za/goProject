/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-20 22:07:05
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-20 22:15:36
 * @FilePath: /src/chapter09/mapstruct/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

type Stu struct {
	Name string
	Age int
	Address string

}
func main()  {
	students := make(map[string]Stu)
	stu1 := Stu{"tom", 20, "shanghai"}
	stu2 := Stu{"jack", 40, "shanghai"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)
	
	for k, v := range students {
		fmt.Println("==============")
		fmt.Println("学号：", k)
		fmt.Println("姓名：", v.Name)
		fmt.Println("年龄：", v.Age)
		fmt.Println("地址：", v.Address)
		fmt.Println("==============")
	}
}