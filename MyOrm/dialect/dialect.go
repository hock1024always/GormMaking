package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

// Dialect is an interface contains methods that a dialect has to implement
type Dialect interface {
	//用于将 Go 语言的类型转换为该数据库的数据类型
	DataTypeOf(typ reflect.Value) string
	//返回某个表是否存在的 SQL 语句，参数是表名(table)
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect 函数用于注册一个方言
func RegisterDialect(name string, dialect Dialect) {
	// 将方言添加到方言映射表中
	dialectsMap[name] = dialect
}

// 根据名称获取方言
func GetDialect(name string) (dialect Dialect, ok bool) {
	// 从方言映射中获取指定名称的方言
	dialect, ok = dialectsMap[name]
	// 返回获取到的方言和是否成功获取的布尔值
	return
}
