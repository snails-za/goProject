/*
* @Author: wangju wangjuchn@outlook.com
* @Date: 2022-12-17 01:31:16
  - @LastEditors: wangju wangjuchn@outlook.com
  - @LastEditTime: 2022-12-17 01:36:30

* @FilePath: /src/struct_tag/struct_mapstructure/main.go
* @Description: 字段标签
默认情况下，mapstructure使用结构体中字段的名称做这个映射，例如我们的结构体有一个Name字段，
mapstructure解码时会在map[string]interface{}中查找键名name。注意，这里的name是大小写不敏感的！
当然，我们也可以指定映射的字段名。为了做到这一点，我们需要为字段设置mapstructure标签。例如下面使用username代替上例中的name：
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string `mapstructure:"username"`
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	datas := []string{`
	  { 
		"type": "person",
		"username":"dj",
		"age":18,
		"job": "programmer"
	  }
	`,
		`
	  {
		"type": "cat",
		"name": "kitty",
		"Age": 1,
		"breed": "Ragdoll"
	  }
	`,
		`
	  {
		"type": "cat",
		"Name": "rooooose",
		"age": 2,
		"breed": "shorthair"
	  }
	`,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		// 我们使用标签mapstructure:"username"将Person的Name字段映射为username，在 JSON 串中我们需要设置username才能正确解析。
		// 另外，注意到，我们将第二个 JSON 串中的Age和第三个 JSON 串中的Name首字母大写了，但是并没有影响解码结果。
		// mapstructure处理字段映射是大小写不敏感的。
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat", cat)
		}
	}
}
