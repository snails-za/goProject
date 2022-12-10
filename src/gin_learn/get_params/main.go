/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 16:57:26
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 17:07:07
 * @FilePath: /gin_learn/get_params/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	router.GET("welcome", welcome)
	router.Run()
}

func welcome(ctx *gin.Context) {
	firstname := ctx.DefaultQuery("firstname", "jack")
	lastname := ctx.Query("lastname")
	ctx.JSON(http.StatusOK, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}
