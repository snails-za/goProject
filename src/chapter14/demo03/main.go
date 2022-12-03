/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-01 23:29:42
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-03 00:58:22
 * @FilePath: /src/chapter14/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sync"
)

func SyncClass() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println(i)
		})
	}
}

func main() {
	// sync.RWMutex{} 读写互斥锁
	// Lock() 和 Unlock()
	// RLock() 和 RUnlock()
	SyncClass()
}
