package main

import "fmt"
func main() {
	/*
	一个养鸡场有6只鸡，他们的体重分别是3kg，5kg，1kg，3.4kg，2kg，50kg。请问这六只鸡的总体重是多少？平均体重是多少？请你编一个程序。=》数组 
	*/

	// 思路分析：定义留个变量，分别表示六只鸡的体重，然后求出和，然后求出平均值。
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0

	totalweight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6  // invalid operation: hen1 + hen2 + hen3 + hen4 (mismatched types int and float64)go
	avgweight := totalweight / 6

	fmt.Printf("totalweight=%.2f \n", totalweight)
	fmt.Printf("totalweight=%v avgweight=%v \n", totalweight, avgweight)

	// 使用数组方式来解决
	// 1.定义一个数组
	var hens [6]float64
	// 给数组的每个元素赋值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	// 遍历数组求出总体重
	totalweight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalweight2 += hens[i]
	}
	avgweight2 := totalweight2 / float64(len(hens))
	fmt.Printf("hens[1:]=%v \n", hens[1:])
	fmt.Printf("totalweight2=%.2f \n", totalweight2)
	fmt.Printf("totalweight2=%v avgweight2=%v \n", totalweight2, avgweight2)

}