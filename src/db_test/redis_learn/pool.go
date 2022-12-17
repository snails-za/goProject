/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-17 23:34:47
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-17 23:41:39
 * @FilePath: /db_test/redis_learn/pool.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package redis_learn

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool //创建redis连接池

func PoolUtil() {
	Pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func PoolTest() {
	c := Pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	_, err := c.Do("Set", "abc", 888)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	fmt.Println(r)
	Pool.Close() //关闭连接池
}
