/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 18:33:06
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 18:42:49
 * @FilePath: /gin_learn/form_params/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func form(ctx *gin.Context) {
	types := ctx.DefaultPostForm("type", "post")
	username := ctx.PostForm("username")
	userpassword := ctx.PostForm("userpassword")
	ctx.JSON(http.StatusOK, gin.H{
		"types":        types,
		"usernmae":     username,
		"userpassword": userpassword,
	})
}

func main() {
	router := gin.Default()
	router.POST("/form", form)
	router.Run()
}
