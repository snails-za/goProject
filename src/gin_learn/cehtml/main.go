/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 19:19:22
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 19:33:18
 * @FilePath: /cehtml/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tem/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "我是测试", "address": "www.5lmh.com"})
	})
	r.Run()
}
