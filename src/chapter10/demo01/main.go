/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-26 22:13:40
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-26 22:24:16
 * @FilePath: /src/chapter10/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"
type Qimiao struct {
	Name string
	Age  int
	Sex  bool
	Hobbys []string
}
func main()  {
	// 方法一
	var qm Qimiao
	qm.Name = "qimiao"
	qm.Age = 10
	qm.Sex = true
	qm.Hobbys = []string{"play", "Song"}
	fmt.Println(qm)
	// 方法二
	qm2 := Qimiao{
		Name: "qimiao",
		Age:  10,
		Sex:  true,
		Hobbys: []string{"play", "sang"},
	}
	fmt.Println(qm2)
	// 方法三
	qm3 := Qimiao{"qimiao", 10, true, []string{"play", "sang"}}
	fmt.Println(qm3)
	// 方法四
	qm4 := new(Qimiao)
	qm4.Name = "qimiao"
	fmt.Println(qm4)
}