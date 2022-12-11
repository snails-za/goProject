/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 18:27:46
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 18:34:51
 * @FilePath: /url_bind/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:user/:password", func(ctx *gin.Context) {
		// 声明接收的变量
		var login Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := ctx.ShouldBindUri(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Pssword != "admin" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run()
}
