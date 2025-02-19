package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		fmt.Println("无法打开数据库:", err)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("无法关闭数据库:", err)
		}
	}()

	_, err = db.Exec("DROP TABLE IF EXISTS User;")
	if err != nil {
		fmt.Println("删除表时出错:", err)
		return
	}

	_, err = db.Exec("CREATE TABLE User(Name text);")
	if err != nil {
		fmt.Println("创建表时出错:", err)
		return
	}

	// 正确的插入语句，使用两次插入
	_, err = db.Exec("INSERT INTO User(`Name`) values (?)", "Tom")
	if err != nil {
		fmt.Println("插入数据时出错:", err)
		return
	}
	result, err := db.Exec("INSERT INTO User(`Name`) values (?)", "Sam")
	if err != nil {
		fmt.Println("插入数据时出错:", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响的行数时出错:", err)
		return
	}
	fmt.Println("受影响的行数:", affected)

	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	err = row.Scan(&name)
	if err != nil {
		fmt.Println("查询数据时出错:", err)
		return
	}
	fmt.Println("查询到的名称:", name)
}
