/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-11-26 22:13:40
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-11-26 23:00:20
 * @FilePath: /src/chapter10/demo01/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package main

 import "fmt"
 type Qimiao struct {
	 Name string
	 Age  int
	 Sex  bool
	 Hobbys []string
	 Home
 }
 type Home struct {
	P string
 }
 
 func (q *Qimiao) Song(name string)(retur string) {
	fmt.Println(q.Name)
	fmt.Println(name)
	retur = "666"
	return retur
}
 
 func (h *Home)Open()  {
	fmt.Println("8888", h.P)
 }
 
 func main()  {
	 var qm Qimiao
	 qm.Name = "qimiao"
	 qm.Age = 10
	 qm.Sex = true
	 qm.Hobbys = []string{"play", "Song"}
	 fmt.Println(qm)
	 restr := qm.Song("凤凰传奇")
	 fmt.Println(restr)
	 qm.Home.P = "nihao"
	 fmt.Println(qm)
	 qm.Home.Open()
 }