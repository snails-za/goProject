/*
  - @Author: wangju wangjuchn@outlook.com
  - @Date: 2022-12-11 21:20:00
  - @LastEditors: wangju wangjuchn@outlook.com
  - @LastEditTime: 2022-12-11 22:55:30
  - @FilePath: /src/gin_learn/cookie/main.go
  - @Description:
    中间件代码最后即使没有调用Next()方法，后续中间件及handlers也会执行；
    如果在中间件函数的非结尾调用Next()方法当前中间件剩余代码会被暂停执行，会先去执行后续中间件及handlers，等这些handlers全部执行完以后程序控制权会回到当前中间件继续执行剩余代码；
    如果想提前中止当前中间件的执行应该使用return退出而不是Next()方法；
    如果想中断剩余中间件及handlers应该使用Abort方法，但需要注意当前中间件的剩余代码会继续执行。
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	// 1.创建路由
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8000")
}
