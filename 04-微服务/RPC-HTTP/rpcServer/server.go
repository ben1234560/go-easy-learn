package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

// golang实现RPC程序，实现求矩形面积和周长

// 1.定义结构体，用于绑定方法
// 矩形对象，首字母大写
type Rect struct {
}

// 2.声明参数的结构体（也可以声明返回值的结构体，因为这里只返回一个值，所以不需要做结构体）
// 首字母大写，参数首字母也大写
type Params struct {
	Width, Height int
}

// 求矩形面积
// (r *Rect)：绑定到矩形对象上，Area：方法名（首字母大写），p:参数对象，Params：参数内容
// ret:返回值对象，*int:返回int类型，参数必须是指针类型，error：必须有error返回声明
func (r *Rect) Area(p Params, ret *int) error {
	// 面积的计算公式，上面声明的是指针，那么赋值的也需要是指针
	*ret = p.Width * p.Height
	// 如上面定义，ret是一定回返回，error是有问题则返回，那么没问题就返回nil，或者你自己定义一个正确时的返回标识
	return nil
}

// 求周长
// 如上解释一样
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	// 1.注册服务，结果不需要所以用 _ 表示抛弃，这里注册的服务必须是所有函数绑定的对象，也就是Rect
	_ = rpc.Register(new(Rect))
	// 2.服务绑定到http协议上
	rpc.HandleHTTP()
	// 3.服务端监听，端口是8000
	err := http.ListenAndServe(":8000", nil)
	// 处理下异常
	if err != nil {
		fmt.Println(err)
		return
	}
	// 这样服务端就写完了
}
