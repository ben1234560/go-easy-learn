package main

import (
	"context"
	"fmt"
	pb "go/gPRC/proto"
	"google.golang.org/grpc"
)

func main()  {
	// 1.创建与gPRC服务端的连接
	conn, e := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if e != nil{
		fmt.Printf("连接异常：%s\n",e)
	}
	// 最后关闭连接，否则消耗资源,defer表示最后执行
	defer conn.Close()
	// 2.实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 3.组装参数，初始化一个指向类型的指针(*T)
	req := new(pb.UserRequest)
	req.Name = "ben"
	// 4.调用接口并传入参数
	response, e := client.GetUserInfo(context.Background(), req)
	if e != nil{
		fmt.Printf("响应异常%s\n", e)
	}
	fmt.Printf("响应结果%v\n", response)
}