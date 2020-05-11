package main

import (
	"context"
	"fmt"
	"github.com/go-micro"
	"github.com/go-micro/errors"
	pb "go/go-micro/micro-http/proto"
	"log"
)

// 定义使用的结构体
type Example struct {
}

type Foo struct {
}

// 编写接口的方法
// 语法：func(绑定结构体对象名 绑定结构体) 方法名(ctx context.Context固定写法 传入参数名及参数)(输出参数名及参数 返回错误){函数体}
func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rep *pb.CallResponse) error {
	log.Println("收到 Example.Call 请求")
	// 如果没有传参
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no context")
	}
	rep.Message = "PRC Call 收到请求" + req.Name
	return nil
}

func (e *Foo) Emp(ctx context.Context, req *pb.EmptyRequest, rep *pb.EmptyResponse) error {
	log.Println("收到 Foo.Emp 请求")
	return nil
}

// 编写主函数
func main() {
	// 实例化service
	service := micro.NewService(
		micro.Name("go.micro.api.index"),
		// 如果报408timeout错误则指定地址，即把下面注释的内容解开
		//micro.Address("127.0.0.1:8080"),
	)
	// 初始化
	service.Init()
	// 注册example接口
	e := pb.RegisterExampleHandler(service.Server(), new(Example))
	if e != nil {
		fmt.Println(e)
	}
	// 注册foo接口
	e2 := pb.RegisterFooHandler(service.Server(), new(Foo))
	if e2 != nil {
		fmt.Println(e2)
	}
	// 运行service并处理可能的报错
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
