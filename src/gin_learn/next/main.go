/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 22:49:03
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 22:49:16
 * @FilePath: /src/gin_learn/next/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	// {}为了代码规范
	{
		r.GET("/ce", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	r.Run()
}
