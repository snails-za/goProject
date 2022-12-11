/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 18:12:33
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 18:25:33
 * @FilePath: /form_bind/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User    string `form"usernmae" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/loginForm", func(ctx *gin.Context) {
		// 声明接收的变量
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := ctx.Bind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "root" || form.Pssword != "admin" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": 200})
	})
	r.Run()
}
