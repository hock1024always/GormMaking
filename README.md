# 轮子项目————使用go语言实现常用数据库的操作

## 项目计划
1. 学习实现Mysql的基础操作，结合实现redis
2. 实现Mysql复杂操作，如事务，索引，存储过程等
3. 实现一些其他相关数据库的操作

## 项目依赖
1. 安装sqlite的dll和tools包，并且解压到一个目录下面（配置环境变量）
2. 依赖不用管，建议直接go mod tidy
3. 安装gcc环境并且配置cgo：$env:CGO_ENABLED="1"
4. 

## 存在问题及解决方法

### 通用数据库操作框架对于不同数据类型的操作。如何根据任意类型的指针，得到其对应的结构体的信息。

涉及到了 Go 语言的反射机制(reflect)，通过反射，可以获取到对象对应的结构体名称，成员变量、方法等信息