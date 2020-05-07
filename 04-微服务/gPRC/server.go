package main

// 这些包再编写代码时会被自动导入，只要确保导入的包是正确的即可，因为有可能有同名包
import (
	"context"
	"fmt"
	pb "go/gPRC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 定义服务端实现约定的接口（结构体形式）
type UserInfoService struct {
}

// 用于注册，下面会用到
var u = UserInfoService{}

// 实现服务端需要实现的接口
// 方法是GetUserInfo，绑定到*UserInfoService，并命名为s，ctx是go语法一般都会有的东西，可能用不上但一定会有
// 传入和返回参数参考pb.go里面的，根据go语法，err是必须返回的
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	// 获取传过来的信息
	name := req.Name
	// 在数据库查用户信息，模拟数据库，判断是否是数据库中的用户
	if name == "ben" {
		resp = &pb.UserResponse{
			Id:   1,
			Name: name,
			Age:  22,
			// 切片字段
			Hobby: []string{"Learn", "Run"},
		}
	}
	err = nil
	return
}

// 程序入口
func main() {
	// 1.监听规则
	addr := "127.0.0.1:8000"
	listener, e := net.Listen("tcp", addr)
	// 处理可能的异常
	if e != nil {
		fmt.Printf("监听异常：%s\n", e)
	}
	// 如果没有异常打印正常信息
	fmt.Printf("开始监听：%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	// 第二个参数要求接口类型的变量
	pb.RegisterUserInfoServiceServer(s, &u)
	// 4.启动gRPC服务端
	e = s.Serve(listener)
	// 处理可能的异常
	if e != nil {
		log.Fatal(e)
	}
}
