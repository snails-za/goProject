package main


import "fmt"
func main() {
	arr := []int{1,2,3,4}
	demo("hello", arr...)
}
func demo(params1 string, params2 ...int)  {
	fmt.Println(params1)
	for _, v := range params2 {
		fmt.Println(v)
	}	
}