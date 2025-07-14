package schema

import (
	"github.com/hock1024always/GormMaking/dialect"
	"go/ast"
	"reflect"
)

// Field represents a column of database
type Field struct {
	Name string // 字段名
	Type string // 字段类型
	Tag  string // 约束条件
}

// Schema
type Schema struct {
	Model      interface{}       //被映射的对象
	Name       string            //表名
	Fields     []*Field          //字段
	FieldNames []string          //包含所有的字段名(列名)
	fieldMap   map[string]*Field //记录字段名和 Field 的映射关系，方便之后直接使用，无需遍历 Fields
}

// GetField returns field by name
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Values return the values of dest's member variables
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

type ITableName interface {
	TableName() string
}

// Parse a struct to a Schema instance
// Parse函数用于解析传入的dest参数，并返回一个Schema结构体
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// 获取dest的反射类型
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	// 定义表名
	var tableName string
	// 判断dest是否实现了ITableName接口
	t, ok := dest.(ITableName)
	if !ok {
		// 如果没有实现，则使用modelType的Name作为表名
		tableName = modelType.Name()
	} else {
		// 如果实现了，则使用ITableName接口的TableName方法作为表名
		tableName = t.TableName()
	}
	// 创建Schema结构体
	schema := &Schema{
		Model:    dest,
		Name:     tableName,
		fieldMap: make(map[string]*Field),
	}

	// 遍历modelType的字段
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 判断字段是否为匿名字段，并且字段名是否为导出的
		if !p.Anonymous && ast.IsExported(p.Name) {
			// 创建Field结构体
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			// 判断字段是否有geeorm标签
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			// 将Field结构体添加到Schema结构体的Fields字段中
			schema.Fields = append(schema.Fields, field)
			// 将字段名添加到Schema结构体的FieldNames字段中
			schema.FieldNames = append(schema.FieldNames, p.Name)
			// 将字段名和Field结构体添加到Schema结构体的fieldMap字段中
			schema.fieldMap[p.Name] = field
		}
	}
	// 返回Schema结构体
	return schema
}
