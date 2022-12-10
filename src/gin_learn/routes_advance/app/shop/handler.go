/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 23:35:16
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 23:36:47
 * @FilePath: /routes_advance/app/shop/handler.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Shop(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "shopping",
		"data": "",
	})
}
