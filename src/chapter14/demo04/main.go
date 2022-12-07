/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-03 13:52:49
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-03 14:00:34
 * @FilePath: /src/chapter14/demo04/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(8 * time.Second)
		wg.Done()
		fmt.Println("hp-1")
	}()
	go func() {
		time.Sleep(6 * time.Second)
		wg.Done()
		fmt.Println("hp-2")
	}()
	wg.Wait()
	fmt.Println("hp=0")
}
