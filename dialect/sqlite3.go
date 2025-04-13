package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

var _ Dialect = (*sqlite3)(nil)

// 初始化函数，用于注册sqlite3方言
func init() {
	// 注册sqlite3方言，参数1为方言名称，参数2为sqlite3方言的结构体
	RegisterDialect("sqlite3", &sqlite3{})
}

// DataTypeOf 函数用于获取 reflect.Value 类型的数据类型
func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	// 根据不同的类型返回不同的数据类型
	// 根据 reflect.Type 的 Kind() 方法返回的类型，映射到对应的 SQLite 数据类型
	switch typ.Kind() {
	// 如果 Kind() 返回的是 Bool，则返回 SQLite 的 bool 类型
	case reflect.Bool:
		return "bool"
		// 如果 Kind() 返回的是 Int, Int8, Int16, Int32, Uint, Uint8, Uint16, Uint32, Uintptr 中的任意一种，则返回 SQLite 的 integer 类型
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
		// 如果 Kind() 返回的是 Int64 或 Uint64，则返回 SQLite 的 bigint 类型
	case reflect.Int64, reflect.Uint64:
		return "bigint"
		// 如果 Kind() 返回的是 Float32 或 Float64，则返回 SQLite 的 real 类型
	case reflect.Float32, reflect.Float64:
		return "real"
		// 如果 Kind() 返回的是 String，则返回 SQLite 的 text 类型
	case reflect.String:
		return "text"
		// 如果 Kind() 返回的是 Array 或 Slice，则返回 SQLite 的 blob 类型
	case reflect.Array, reflect.Slice:
		return "blob"
		// 如果 Kind() 返回的是 Struct，则需要进一步判断
	case reflect.Struct:
		// 如果 Struct 是 time.Time 类型，则返回 SQLite 的 datetime 类型
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}

	// 如果类型不符合要求，则抛出异常
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

// 判断表是否存在
func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	// 定义参数
	args := []interface{}{tableName}
	// 返回查询语句和参数
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
