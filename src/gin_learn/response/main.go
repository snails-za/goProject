/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-11 18:42:29
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-11 18:54:25
 * @FilePath: /response/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func json_resp(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "someJSON", "status": 200})
}

func struct_resp(ctx *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}
	msg.Name = "root"
	msg.Message = "message"
	msg.Number = 123
	ctx.JSON(200, msg)
}

func xml_resp(ctx *gin.Context) {
	ctx.XML(200, gin.H{"message": "abc"})
}
func yaml_resp(c *gin.Context) {
	c.YAML(200, gin.H{"name": "zhangsan"})
}
func protobuf_resp(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	// 定义数据
	label := "label"
	// 传protobuf格式数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(200, data)
}

// 多种响应方式
func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()

	// 1.json
	r.GET("/someJSON", json_resp)

	// 2. 结构体响应
	r.GET("/someStruct", struct_resp)

	// 3.XML
	r.GET("someXml", xml_resp)

	// 4.YAML响应
	r.GET("/someYAML", yaml_resp)

	// 5.protobuf格式,谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	r.GET("/someProtoBuf", protobuf_resp)
	r.Run()
}
