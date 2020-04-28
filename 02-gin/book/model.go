package main

// 建结构体（可以理解成数据模板），规定数据类型用于被继承，根据数据库（所需字段）来定义
// 结构体名称必须大写开头，否则json转换会产生问题，而字段也必须大写开头，否则其它文件无法引用
type Book struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
	Price int64  `db:"price"`
}
