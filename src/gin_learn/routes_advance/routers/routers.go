/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 23:39:03
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 23:47:18
 * @FilePath: /routes_advance/routers/routers.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

func Include(ots ...Option) {
	options = append(options, ots...)
}

func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}
