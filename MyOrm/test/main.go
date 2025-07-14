package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //匿名导入,这种导入方式的主要目的是为了执行包中的 init 函数
)

// 这次程序用来测试使用go语言操作sqlite3数据库的基本操作 增删改查
func main() {
	// 打开数据库
	db, _ := sql.Open("sqlite3", "gee.db")
	// 关闭数据库
	defer func() { _ = db.Close() }()
	// 删除User表
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	// 创建User表
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	// 向User表中插入数据
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	// 如果插入数据成功
	if err == nil {
		// 获取受影响的行数
		affected, _ := result.RowsAffected()
		// 打印受影响的行数
		log.Println(affected)
	}
	// 查询User表中的第一条数据
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	// 定义变量name
	var name string
	// 如果查询成功
	if err := row.Scan(&name); err == nil {
		// 打印name
		log.Println(name)
	}
}
