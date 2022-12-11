/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 20:37:45
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 20:39:28
 * @FilePath: /redirect/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.5lmh.com")
	})
	r.Run()
}
