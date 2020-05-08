package main

import (
	"fmt"
	"net/rpc"
)

// 客户端需要传这些参数，也需要这个结构体
type Params struct {
	Width, Height int
}

func main() {
	// 1.连接RPC远程服务
	// rpc也是基于tcp做的，端口对应上服务端的端口
	client, err1 := rpc.DialHTTP("tcp", ":8000")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	// 2.调用远程方法
	// 定义一个结果接收对象
	ret := 0
	// Rect.Area：调用面积计算函数，Params：传给函数的参数，&ret：返回值（用地址接收，因为那边是指针）
	err2 := client.Call("Rect.Area", Params{10, 20}, &ret)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("面积：", ret)
	// 调用周长计算函数
	err3 := client.Call("Rect.Perimeter", Params{10, 20}, &ret)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println("周长：", ret)
}