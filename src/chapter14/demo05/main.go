/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-03 14:10:20
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-03 14:21:11
 * @FilePath: /src/chapter14/demo05/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := &sync.Map{}
	m.Store(1, 1)
	fmt.Println(m.Load(1))
	fmt.Println(m.Load(1))
	m.Delete(1)
	fmt.Println(m.Load(1))
	m.Store(1, 1)
	m.Store(2, 2)
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	time.Sleep(5 * time.Second)
}
