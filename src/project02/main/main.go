package main

import (
	"fmt"
	"project02/model"
)

func main() {
	fmt.Println(model.Heroname)
	model.DbConnect()
}