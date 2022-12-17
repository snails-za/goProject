/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-17 22:44:53
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-17 23:31:00
 * @FilePath: /db_test/redis_learn/string.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package redis_learn

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func Set(c redis.Conn, key string, value interface{}) {

	_, err := c.Do("Set", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func Get(c redis.Conn, key string) {
	r, err := redis.Int(c.Do("Get", key))
	if err != nil {
		fmt.Println("get "+key+" failed,", err)
		return
	}

	fmt.Println(r)
}
func MSet(c redis.Conn) {
	_, err := c.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func MGet(c redis.Conn) {
	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	for _, v := range r {
		fmt.Println(v)
	}
}
func Expire(c redis.Conn, key string, ex int) {
	_, err := c.Do("expire", key, ex)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func Lpush(c redis.Conn) {
	_, err := c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func Lpop(c redis.Conn, key string) {
	r, err := redis.String(c.Do("lpop", key))
	if err != nil {
		fmt.Println("get "+key+" failed,", err)
		return
	}

	fmt.Println(r)
}
func HSet(c redis.Conn) {
	_, err := c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func HGet(c redis.Conn) {
	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}
