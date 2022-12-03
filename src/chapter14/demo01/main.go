/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-01 23:29:42
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-03 00:30:17
 * @FilePath: /src/chapter14/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncClass() {
	l := &sync.Mutex{}
	go lockFunc(l)
	go lockFunc(l)
	go lockFunc(l)
	go lockFunc(l)
	time.Sleep(5 * time.Second)
}
func lockFunc(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("666")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}
func main() {
	// sync.Mutex{} 互斥锁
	// Lock() 和 Unlock()
	SyncClass()
}
