
1. `_ "github.com/mattn/go-sqlite3"`
   - 匿名导入,这种导入方式的主要目的是为了执行包中的 init 函数
   - 通过空白导入这个包，可以将 SQLite3 驱动注册到 database/sql 包中，这样你就可以在后续的代码中使用 database/sql 的函数来操作 SQLite3 数据库，而不需要直接调用 github.com/mattn/go-sqlite3 包中的任何函数或变量。