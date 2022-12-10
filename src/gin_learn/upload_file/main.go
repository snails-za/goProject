/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-10 18:52:46
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-10 19:11:24
 * @FilePath: /gin_learn/upload_file/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func upload_file(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败！",
			"data": "",
		})
	}
	ctx.SaveUploadedFile(file, file.Filename)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功！",
		"data": "",
	})

}

func upload_file_size(ctx *gin.Context) {
	_, headers, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败！",
			"data": "",
		})
	}
	//headers.Size 获取文件大小
	if headers.Size > 1024*1024*2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件太大！",
			"data": "",
		})
		return
	}
	//headers.Header.Get("Content-Type")获取上传文件的类型
	if headers.Header.Get("Content-Type") != "image/png" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "只允许上次png",
			"data": "",
		})
		return
	}
	ctx.SaveUploadedFile(headers, headers.Filename)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功！",
		"data": "",
	})

}
func main() {
	router := gin.Default()
	router.POST("upload", upload_file)
	router.POST("upload_size", upload_file_size)
	router.Run()
}
