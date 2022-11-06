package main

import "fmt"
func main() {
	num1 := 10
	num2 := 20
	res1 := max(num1, num2)
	fmt.Println(res1)

	str1 := "google"
	str2 := "runoob"
	res2, res3 := swap(str1, str2)
	fmt.Println(res2, res3)
}


func max(num1, num2 int) int {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}


func swap(a, b string) (string, string) {
	return a, b
}
