package main

import (
	"context"
	"fmt"
	"github.com/go-micro"
	pb "go/go-micro/HelloWorld/proto"
	"log"
)

// 声明结构体
type Hello struct {
}

// 实现接口方法，就是在proto写的接口
// *Hello：把方法绑定到它，Info：名字叫它，ctx context.Context：固定写法
// *pb.InfoRequest：自动实现的传入参数，req：传入对象，*pb.InfoResponse：返回参数，rep：返回对象
func (g *Hello) Info(ctx context.Context, rep *pb.InfoRequest, req *pb.InfoResponse) error{
	// 返回值是 你好 加上 传入参数
	req.Msg = "你好"+rep.Username
	// err还要返回，所以直接写一个return nil
	return nil
}

// 主函数
func main()  {
	// 1.得到微服务实例
	service := micro.NewService(
		// 设置微服务的名字，用来访问，到时候访问hello就能访问到
		micro.Name("hello"),
		// 如果报408timeout错误则指定地址，即把下面注释的内容解开
		// micro.Address("127.0.0.1:54782"),
		)
	// 2.初始化
	service.Init()
	// 3.服务注册
	e := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if e != nil{
		fmt.Println(e)
	}
	// 4.启动微服务，下面的写法是上面的简化版
	if e := service.Run(); e != nil{
		log.Fatal(e)
	}
}
// 写完就可以右键run这个server，然后在cmd访问