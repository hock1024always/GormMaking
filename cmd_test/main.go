package main

import (
	"fmt"
	"github.com/hock1024always/GormMaking"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 创建一个sqlite3数据库引擎，并指定数据库名为gee.db
	engine, _ := GormMaking.NewEngine("sqlite3", "gee.db")
	// 程序结束时关闭数据库引擎
	defer engine.Close()
	// 创建一个新的会话
	s := engine.NewSession()
	// 执行一条SQL语句，删除名为User的表
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	// 执行一条SQL语句，创建名为User的表，表中有一个text类型的Name字段
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	// 再次执行一条SQL语句，创建名为User的表，表中有一个text类型的Name字段
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	// 执行一条SQL语句，向User表中插入两条数据，Name字段分别为Tom和Sam
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	// 获取执行结果，获取受影响的行数
	count, _ := result.RowsAffected()
	// 打印执行结果
	fmt.Printf("Exec success, %d affected\n", count)

}
