/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-17 19:32:36
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-17 22:25:30
 * @FilePath: /db_test/mysql_learn/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
CREATE TABLE `person` (

	  `user_id` int(11) NOT NULL AUTO_INCREMENT,
	  `username` varchar(260) DEFAULT NULL,
	  `sex` varchar(260) DEFAULT NULL,
	  `email` varchar(260) DEFAULT NULL,
	  PRIMARY KEY (`user_id`)
	) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE place (

	country varchar(200),
	city varchar(200),
	telcode int

)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
*/
package mysql_learn

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}
type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func Insert(Db *sqlx.DB) {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func Select(Db *sqlx.DB) {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select succ:", person)
}

func Update(Db *sqlx.DB) {
	res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}

func Delete(Db *sqlx.DB) {
	res, err := Db.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete succ: ", row)
}
