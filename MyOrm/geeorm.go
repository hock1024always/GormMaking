package GormMaking

import (
	"database/sql"
	"github.com/hock1024always/GormMaking/dialect"

	"github.com/hock1024always/GormMaking/log"
	"github.com/hock1024always/GormMaking/session"
)

// Engine
//
//	type Engine struct {
//		db *sql.DB
//	}
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// NewEngine 函数用于创建一个新的数据库引擎
func NewEngine(driver, source string) (e *Engine, err error) {
	// 使用给定的驱动和源打开数据库连接
	db, err := sql.Open(driver, source)
	if err != nil {
		// 如果打开数据库连接失败，记录错误并返回
		log.Error(err)
		return
	}
	// 测试和数据库的连接是否存在
	//if err = db.Ping(); err != nil {
	//	//log.Error(err)
	//	log.Errorf("dialect %s Not Found", driver)
	//	return
	//}
	dial, ok := dialect.GetDialect(driver) //根据驱动获取数据库方言
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	//e = &Engine{db: db}
	e = &Engine{db: db, dialect: dial} //数据库连接成功，返回Engine对象
	log.Info("成功连接数据库")
	return
}

// 关闭数据库的连接
// 关闭数据库连接
func (engine *Engine) Close() {
	// 关闭数据库连接
	if err := engine.db.Close(); err != nil {
		log.Error("关闭数据库连接失败")
	}
	log.Info("关闭数据库连接成功")
}

// // NewSession 方法用于创建一个新的 session
// func (engine *Engine) NewSession() *session.Session {
//
//		// 返回一个新的 session，使用 engine 的 db
//		return session.New(engine.db)
//	}
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
