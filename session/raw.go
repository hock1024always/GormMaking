package session

import (
	"database/sql"
	"github.com/hock1024always/GormMaking/log"
	"strings"
)

// 下面是核心组件Session，负责与数据库的交互

// 创建的这个组件 实现了对数据库的连接、参数记录、语句拼接 以实现数据库的操作
type Session struct {
	db      *sql.DB         // 创建的数据库连接
	sql     strings.Builder //拼接字符串
	sqlVars []interface{}   // sqlVars 是一个可变数组，用于存储 sql 语句中的参数
}

// 创建Session对象
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// 重启 清除Session中的sql和sqlVars
func (s *Session) Clear() {
	s.sql.Reset()   // 重置sql
	s.sqlVars = nil // 将sqlVars置为nil
}

// 返回数据库连接
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw方法用于向Session中添加原始的SQL语句和参数
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	// 将SQL语句添加到Session的sql字段中
	s.sql.WriteString(sql)
	// 在SQL语句后添加一个空格
	s.sql.WriteString(" ")
	// 将参数添加到Session的sqlVars数组中
	s.sqlVars = append(s.sqlVars, values...)
	// 返回Session对象
	return s
}

// Exec函数用于执行SQL语句，并返回执行结果和错误信息
func (s *Session) Exec() (result sql.Result, err error) {
	// 在函数结束时调用Clear函数，清除Session中的数据
	defer s.Clear()
	// 打印SQL语句和参数
	log.Info(s.sql.String(), s.sqlVars)
	// 执行SQL语句，并返回执行结果和错误信息
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow 方法用于执行查询并返回一行结果
func (s *Session) QueryRow() *sql.Row {
	// 在函数结束时调用 Clear 方法
	defer s.Clear()
	// 打印查询语句和参数
	log.Info(s.sql.String(), s.sqlVars)
	// 执行查询并返回一行结果
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows函数用于执行SQL查询并返回结果集
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	// 在函数结束时调用Clear函数
	defer s.Clear()
	// 打印SQL语句和参数
	log.Info(s.sql.String(), s.sqlVars)
	// 执行SQL查询并返回结果集和错误信息
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
