package main

import (
	"fmt"
	"project02/model"
)
// 如果导报失败，请采用以下命令：
// 1. go mod init project02
// 2. go mod tidy
func main() {
	fmt.Println(model.Heroname)
	model.DbConnect()
}