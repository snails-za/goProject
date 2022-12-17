/*
 * @Author: wangju wangjuchn@outlook.com
 * @Date: 2022-12-17 22:29:57
 * @LastEditors: wangju wangjuchn@outlook.com
 * @LastEditTime: 2022-12-17 22:38:58
 * @FilePath: /db_test/mysql_learn/Advance.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mysql_learn

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Advance(Db *sqlx.DB) {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
	}
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}
