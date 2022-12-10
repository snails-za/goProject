/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 23:37:37
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 23:38:29
 * @FilePath: /routes_advance/app/blog/handler.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Blog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "blog",
		"data": "",
	})
}
