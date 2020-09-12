package main

import "fmt"
func main() {
	var a byte
	fmt.Println("请输入字符：")
	fmt.Scanf("%c", &a)
	switch a {
		case 'a':
			fmt.Println("周一")
		case 'b':
			fmt.Println("周二")
		case 'c':
			fmt.Println("周三")
		case 'd', 'e':
			fmt.Println("周四")
		default:
			fmt.Println("有误")
	}

	switch age := 10; {
		case age == 10:
			fmt.Println("age == 10")
		case age == 20:
			fmt.Println("age == 20")
		case age == 30:
			fmt.Println("age == 30")
		default:
			fmt.Println("没有匹配到")

	}

	switch {
	case false:
			fmt.Println("1、case 条件语句为 false")
			fallthrough
	case true:
			fmt.Println("2、case 条件语句为 true")
			fallthrough
	case false:
			fmt.Println("3、case 条件语句为 false")
			fallthrough
	case true:
			fmt.Println("4、case 条件语句为 true")
	case false:
			fmt.Println("5、case 条件语句为 false")
			fallthrough
	default:
			fmt.Println("6、默认 case")
	}

	c := 'A'
	c += 32
	fmt.Print(c)
}