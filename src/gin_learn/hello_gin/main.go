/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-09 22:56:08
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 16:16:12
 * @FilePath: /gin_learn/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "github.com/gin-gonic/gin"

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run()
}
