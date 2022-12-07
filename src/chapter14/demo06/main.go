/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-03 14:10:20
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-03 14:24:41
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
	p := &sync.Pool{}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	time.Sleep(5 * time.Second)
}
