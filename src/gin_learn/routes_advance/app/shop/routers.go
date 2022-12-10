/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 23:32:59
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 23:39:36
 * @FilePath: /routes_advance/app/shop/routes.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package shop

import "github.com/gin-gonic/gin"

func Routers(r *gin.Engine) {
	r.GET("/shop", Shop)
}
