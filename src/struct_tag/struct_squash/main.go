/*
* @Author: wangju wangjuchn@outlook.com
* @Date: 2022-12-17 01:50:56
  - @LastEditors: wangju wangjuchn@outlook.com
  - @LastEditTime: 2022-12-17 01:55:15

* @FilePath: /src/struct_tag/struct_ squash/main.go
* @Description: 内嵌结构
设置mapstructure:",squash"将该结构体的字段提到父结构中
另外需要注意一点，如果父结构体中有同名的字段，那么mapstructure会将JSON 中对应的值同时设置到这两个字段中，即这两个字段有相同的值。
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
}

type Friend1 struct {
	Person
}

type Friend2 struct {
	Person `mapstructure:",squash"`
}

func main() {
	datas := []string{`
	  { 
		"type": "friend1",
		"person": {
		  "name":"dj"
		}
	  }
	`,
		`
	  {
		"type": "friend2",
		"name": "dj2"
	  }
	`,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Println("friend1", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Println("friend2", f2)
		}
	}
}
