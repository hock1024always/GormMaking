package geeorm

import (
	"database/sql"

	"github.com/hock1024always/GormMaking/log"
	"github.com/hock1024always/GormMaking/session"
)

// Engine is the main struct of geeorm, manages all db sessions and transactions.
type Engine struct {
	db *sql.DB
}

// NewEngine create a instance of Engine
// connect database and ping it to test whether it's alive
// NewEngine 函数用于创建一个新的数据库引擎
func NewEngine(driver, source string) (e *Engine, err error) {
	// 使用给定的驱动和源打开数据库连接
	db, err := sql.Open(driver, source)
	if err != nil {
		// 如果打开数据库连接失败，记录错误并返回
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

// Close database connection
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// NewSession creates a new session for next operations
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
