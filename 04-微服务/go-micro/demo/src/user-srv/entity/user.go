package entity

import "go/go-micro/demo/src/share/pb"

// user
type User struct {
	Id      int32  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"phone"`
}

// 定义用于返回的结构体
func (u User) ToProtoUser() *pb.User {
	return &pb.User{
		Id:      u.Id,
		Name:    u.Name,
		Address: u.Address,
		Phone:   u.Phone,
	}
}
