package main

// 这个db包主要是做堆数据库增删改查的操作
// 导入相关的包
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 初始化的db
// 由于其它地方也用到，所以声明一个全局的db，指针类型
var db *sqlx.DB
// 初始化数据库
// 定义函数，func 函数名(传入参数)(返回值){函数体}
func initDB() (err error) {
	// 连接数据库，数据库是mysql，用户名是root，密码是123456，tcp模型（这个不用管），端口号，对应的mysql的库是go_book
	db, err = sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/go")
	// 如果有报错，返回报错
	if err != nil{
		return err
	}
	// 否则返回，这里的返回并没有实际返回什么，但是db已经改变
	return
}

// 功能：查询数据库全部数据，全部数据则不需要传数据，如果是特定数据才需要传id或者其它所需字段
// bookList是返回的值的名字，[]*Book是调用了结构体且是切片类型，因为数据是多行的每个id都有一行，
// 使用结构体必须用*（指针）才能接受到传过来的地址改变结构体的存储值，这样就能在其它地方用到结构体里面的值
func queryAllBook() (bookList []*Book,err error) {
	// SQL语句：从book表查询所有数据，并返回id,title,price（注意不要图省事写*）
	// 必须用&(传地址)传，这里就用到了全局的db
	err = db.Select(&bookList, "SELECT id,title,price FROM book")
	// 这里的返回是指返回bookList，因为函数开头写明了，所以不需要这样写return bookList
	return
}

// 功能：添加数据到数据库
// 传入title和price，id是自增所以不需要传，只需要返回err即可
func insertBook(title string,price int64)(err error){
	// SQL语句，给title，price添加值，值就是传入的title和price
	// 返回的result表示影响行数，不需要所以用 _ 表示丢弃
	_, err = db.Exec("INSERT INTO book(title, price) VALUES(?,?)", title, price)
	return
}

// 功能：删除数据，根据id删除数据
func deleteBook(id int64) (err error) {
	_, err = db.Exec("DELETE FROM BOOK WHERE id=?", id)
	if err != nil{
		return err
	}
	return
}