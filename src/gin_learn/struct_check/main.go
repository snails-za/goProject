/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 23:40:26
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 23:40:45
 * @FilePath: /src/gin_learn/struct_check/main.go
 * @Description: 用gin框架的数据验证，可以不用解析数据，减少if else，会简洁许多。
 */
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Person ..
type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/5lmh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	r.Run()
}
