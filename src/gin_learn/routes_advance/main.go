/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 23:29:30
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 00:07:15
 * @FilePath: /routes_advance/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"routes_advance/app/blog"
	"routes_advance/app/shop"
	"routes_advance/routers"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
