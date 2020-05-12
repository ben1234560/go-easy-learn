package handler

import (
	"go.uber.org/zap"
	"go/go-micro/demo/src/share/pb"
	tlog "go/go-micro/demo/src/share/utils/log"
	"go/go-micro/demo/src/user-srv/db"
	"go/go-micro/demo/src/user-srv/entity"
	"golang.org/x/net/context"
	"log"
)

// 定义绑定方法的结构体
type UserHandler struct {
	logger *zap.Logger
}

// 创建结构体对象
func NewUserHandler() *UserHandler {
	return &UserHandler{
		logger: tlog.Instance().Named("UserHandler"),
	}
}

// 添加
func (c *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserReq, rep *pb.InsertUserRep) error {
	log.Println("InsertUser...")
	// 封装结构体
	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}
	// 调用数据库方法插入
	insertUser, e := db.InsertUser(user)
	if e != nil {
		log.Fatal("添加用户错误")
		return e
	}
	rep.Id = int32(insertUser)
	return nil
}

func (c *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserReq, rep *pb.DeleteUserRep) error {
	log.Println("DeleteUser...")
	e := db.DeleteUser(req.GetId())
	if e != nil {
		log.Fatal("删除用户错误")
		return e
	}
	return nil
}

func (c *UserHandler) ModifyUser(ctx context.Context, req *pb.ModifyUserReq, rep *pb.ModifyUserRep) error {
	log.Println("ModifyUser...", req.GetId())
	// 封装结构体
	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Id:      req.Id,
	}
	e := db.ModifyUser(user)
	if e != nil {
		log.Fatal("修改用户错误")
		return e
	}
	return nil
}

func (c *UserHandler) SelectUser(ctx context.Context, req *pb.SelectUserReq, rep *pb.SelectUserRep) error {
	log.Println("SelectUser...")
	user, e := db.SelectUserById(req.GetId())
	if e != nil {
		log.Fatal("查询用户错误")
		return e
	}
	if user != nil {
		rep.Users = user.ToProtoUser()
	}
	return nil

}
