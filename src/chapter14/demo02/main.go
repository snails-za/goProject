/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-01 23:29:42
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-03 00:51:50
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
	l := &sync.RWMutex{}
	go lockFunc(l)
	go lockFunc(l)
	go lockFunc(l)
	go lockFunc(l)
	go readLockFunc(l)
	go readLockFunc(l)
	go readLockFunc(l)
	go readLockFunc(l)
	time.Sleep(20 * time.Second)
}

func lockFunc(lock *sync.RWMutex) {
	lock.Lock() // 写的时候排斥其他写锁和读锁
	fmt.Println("666")
	time.Sleep(5 * time.Second)
	lock.Unlock()
}

func readLockFunc(lock *sync.RWMutex) {
	lock.RLock() // 在读取的时候不会阻塞其他的读取锁，但是会阻塞写入锁
	fmt.Println("777")
	time.Sleep(5 * time.Second)
	lock.RUnlock()
}

func main() {
	// sync.RWMutex{} 读写互斥锁
	// Lock() 和 Unlock()
	// RLock() 和 RUnlock()
	SyncClass()
}
