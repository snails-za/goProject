/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-13 23:27:29
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-14 23:18:04
 * @FilePath: /src/chapter06/timefunc/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"time"
)
func main()  {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 格式化
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	fmt.Println(now.Format(layout))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}