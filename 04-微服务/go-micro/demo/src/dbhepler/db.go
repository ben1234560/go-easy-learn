package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go/go-micro/demo/src/share/config"
	"jmoiron/sqlx"
	"log"
	"math/rand"
	"time"
)

// 建表语句
var schema = `
CREATE TABLE IF NOT EXISTS user(
	id INT UNSIGNED AUTO_INCREMENT,
	name VARCHAR(20),
	address VARCHAR(20),
	phone VARCHAR(15),
	PRIMARY EKY (id)
)
`

// 对应表的结构体
type User struct {
	Id int32 `db:"id"`
	Name string `db:"name"`
	Address string `db:"address"`
	Phone string `db:"phone"`
}

func main()  {
	// 打开并连接数据库
	db, e := sqlx.Connect("mysql", config.MysqlDNS)
	if e != nil{
		log.Fatal(e)
	}
	// 执行建表语句
	db.MustExec(schema)
	// 添加假数据
	tx := db.MustBegin()
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 事务执行sql插入
	tx.MustExec("INSERT INTO user(id,name,address,phone) VALUES (?,?,?,?)",
		nil, GetRandomString(10),"guangzhou"+GetRandomString(10),
		"1351"+GetRandomString(7))
	err := tx.Commit()
	if err != nil{
		_ = tx.Rollback()
	}
	fmt.Println("执行完毕")
}

// 按指定个数生成随机数
func GetRandomString(leng int) string {
	str := "0123456789abcdefghigklmn"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<leng;i++{
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}