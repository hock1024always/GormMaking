# 轮子项目————使用go语言实现常用数据库的操作

## 项目计划
1. 学习实现Mysql的基础操作，结合实现redis
2. 实现Mysql复杂操作，如事务，索引，存储过程等
3. 实现一些其他相关数据库的操作

## 项目依赖
1. 安装sqlite的dll和tools包，并且解压到一个目录下面（配置环境变量）
2. 依赖不用管，建议直接go mod tidy
3. 安装gcc环境并且配置cgo：$env:CGO_ENABLED="1"

# Day1
## SQLite3 驱动`_ "github.com/mattn/go-sqlite3"`
- 匿名导入,这种导入方式的主要目的是为了执行包中的 init 函数
- 通过空白导入这个包，可以将 SQLite3 驱动注册到 database/sql 包中，这样你就可以在后续的代码中使用 database/sql 的函数来操作 SQLite3 数据库，而不需要直接调用 github.com/mattn/go-sqlite3 包中的任何函数或变量。

## 日志处理
在日志的设置中存在一个互斥锁，来保证日志工具（等级设置）同时只能有一个goroutine进行使用
1. 竞态条件（Race Condition）：多个 goroutine 可能会同时检查 ErrorLevel 和 InfoLevel 的值，并尝试设置日志输出。这会导致不可预测的行为，因为每个 goroutine 可能会看到 ErrorLevel 和 InfoLevel 的不同值，从而做出不同的决策。
2. 数据不一致：如果一个 goroutine 正在修改 loggers 切片或日志记录器的状态，而另一个 goroutine 在同一时间尝试读取这些状态，可能会读取到不完整或不一致的数据。
3. 资源竞争：多个 goroutine 同时写入同一个文件或资源（例如 os.Stdout）可能会导致数据混乱或资源损坏。

## Session
这个结构是负责处理和数据库的交互工作
1. session的创建，是为了复用数据库连接，减少数据库连接的开销，提高性能。
2. 创建session对象之后，有一个clear方法，在每次操作的最后执行，将session对象中的数据清空，以便下次复用。
3. 字段的用途
    - `db *sql.DB`：数据库连接对象，用于执行 SQL 查询。
    - `sql strings.Builder`: 这是一个 strings.Builder 类型的字段，用于高效地拼接 SQL 语句字符串。strings.Builder 是 Go 语言中用于构建字符串的高效工具，特别适合在需要动态构建 SQL 语句的场景中使用。
    - `sqlVars []interface{}`: 这是一个可变数组（切片），用于存储 SQL 语句中的参数。通过将参数存储在 sqlVars 中，可以避免直接将参数拼接到 SQL 字符串中带来的 SQL 注入风险，并且可以使用数据库驱动中的占位符来安全地插入参数。

## Engine
这个结构是负责处理和数据库交互前的准备工作（比如连接/测试数据库），交互后的收尾工作（关闭连接）等
1. 创建Engine，查看和数据库是否正常的连接，如果连接正常，则创建Engine对象
2. 给一个关闭数据库的方法以及一个创建session的方法

# Day2
1. 为适配不同的数据库，映射数据类型和特定的 SQL 语句，创建 Dialect 层屏蔽数据库差异。
2. 设计 Schema，利用反射(reflect)获取任意 struct 对象的名称和字段，映射为数据中的表，包括表名、字段名、字段类型、字段 tag 等。
3. 构造创建(create)、删除(drop)、存在性(table exists) 的 SQL 语句完成数据库表的基本操作。

## Dialect
1. sql语句中，数据类型和golang存在区别，比如整形：golang是int，sql是integer，需要进行转换。Dialect是用来解决这个问题的。
2. 不同数据库支持的数据类型也是有差异的，即使功能相同，在 SQL 语句的表达上也可能有差异
上面两种差异的提取过程就是dialect的作用
## Scheme
对象(object)和表(table)的转换。给定一个任意的对象，转换为关系型数据库中的表结构。
创建一张表需要哪些要素呢？
- 表名(table name) —— 结构体名(struct name)
- 字段名和字段类型 —— 成员变量和类型。
- 额外的约束条件(例如非空、主键等) —— 成员变量的Tag（Go 语言通过 Tag 实现，Java、Python 等语言通过注解实现）

## session的改动
对于原始的session进行了如下改动
1. Session 成员变量新增 dialect 和 refTable
2. 构造函数 New 的参数改为 2 个，db 和 dialect。
同时新增了一个数据库表的功能
1. Model() 方法用于给 refTable 赋值。解析操作是比较耗时的，因此将解析的结果保存在成员变量 refTable 中，即使 Model() 被调用多次，如果传入的结构体名称不发生变化，则不会更新 refTable 的值。
2. RefTable() 方法返回 refTable 的值，如果 refTable 未被赋值，则打印错误日志。

## Engine的改动
1. NewEngine 创建 Engine 实例时，获取 driver 对应的 dialect。
2. NewSession 创建 Session 实例时，传递 dialect 给构造函数 New。

