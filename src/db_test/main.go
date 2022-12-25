/*
* @Author: wangju wangjuchn@outlook.com
* @Date: 2022-12-17 19:31:37
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-25 21:40:19

* @FilePath: /db_test/main.go
* @Description:
数据库操作
*/
package main

import (
	"db_test/conf"
	"db_test/gorm_learn"
	"db_test/mysql_learn"
	_ "db_test/mysql_learn"
	"db_test/redis_learn"
	_ "db_test/redis_learn"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

func main() {

	/*
		MYSQL
	*/
	// 使用第三方开源的mysql库: github.com/go-sql-driver/mysql （mysql驱动） github.com/jmoiron/sqlx （基于mysql驱动的封装）
	var Db *sqlx.DB
	database, err := sqlx.Open("mysql", "root:"+conf.PASSWORD+"@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	defer Db.Close() // 注意这行代码要写在上面err判断的下面
	// mysql_learn.Insert(Db)
	mysql_learn.Select(Db)
	// mysql_learn.Update(Db)
	// mysql_learn.Delete(Db)
	// mysql_learn.Advance(Db)

	/*
		REDIS
	*/
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	// 1、String类型 set、get
	// redis_learn.Set(c, "abc", 20)
	redis_learn.Get(c, "abc")
	// // 2、String类型批量操作
	// redis_learn.MSet(c)
	// redis_learn.MGet(c)
	// // 3、设置过期时间
	// redis_learn.Set(c, "abc", 50)
	// redis_learn.Expire(c, "abc", 10)
	// // time.Sleep(10 * time.Second)
	// time.Sleep(2 * time.Second)
	// redis_learn.Get(c, "abc")
	// // 4、List队列操作
	// redis_learn.Lpush(c)
	// redis_learn.Lpop(c, "book_list")
	// redis_learn.Lpop(c, "book_list")
	// redis_learn.Lpop(c, "book_list")
	// redis_learn.Lpop(c, "book_list")
	// // 5、Hash表
	// redis_learn.HSet(c)
	// redis_learn.HGet(c)
	// // 6、连接池
	// redis_learn.PoolUtil()
	// redis_learn.PoolTest()

	/*
		GORM
	*/
	gorm_learn.QuickStart()
}
