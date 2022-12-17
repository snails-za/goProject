/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-17 00:50:55
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-17 00:54:22
 * @FilePath: /src/struct_tag/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}
	var s2 Student
	json.Unmarshal([]byte(data), &s2)
	fmt.Println(s2)
}
