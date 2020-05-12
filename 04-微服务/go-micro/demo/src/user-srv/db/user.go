package db

import (
	"database/sql"
	"go/go-micro/demo/src/user-srv/entity"
)

// user表的查询
func SelectUserById(id int32) (*entity.User, error) {
	// 创建user对象，用于返回
	user := new(entity.User)
	// 执行查询
	if e := db.Get(user, "SELECT name,address,phone FROM user WHERE id=?", id); e != nil {
		if e != sql.ErrNoRows {
			return nil, e
		}
		return nil, nil
	}
	return user, nil
}

// 增加
func InsertUser(user *entity.User) (int64,error) {
	result, e := db.Exec("INSERT INTO `user`(`name`,`address`,`phone`)VALUES(?,?,?)",
		user.Name, user.Address, user.Phone)
	if e !=nil{
		return 0,e
	}
	return result.LastInsertId()
}

// 修改
func ModifyUser(user *entity.User) (error) {
	_, e := db.Exec("UPDATE `user` set `name`=?,`address`=?,`phone`=? WHERE `id`=?",
		user.Name, user.Address, user.Phone, user.Id)
	if e !=nil{
		return e
	}
	return nil
}
// 删除
func DeleteUser(id int32) error {
	_, e := db.Exec("DELETE FROM `user` WHERE `id`=?",
		id)
	if e !=nil{
		return e
	}
	return nil
}