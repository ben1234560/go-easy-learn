package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go/go-micro/demo/src/share/config"
	"go/go-micro/demo/src/share/pb"
	"go/go-micro/demo/src/share/utils/log"
	"go/go-micro/demo/src/user-srv/db"
	"go/go-micro/demo/src/user-srv/handler"
)

func main() {
	logger := log.Init("user")
	// 1.创建service
	service := micro.NewService(
		micro.Name(config.Namespace+"user"),
		micro.Version("latest"),
		// 如果报408timeout错误则指定地址，即把下面注释的内容解开
		//micro.Address("127.0.0.1:8888"),
	)
	// 2.初始化service
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("user-srv服务运行时的打印...")
			// 初始化db
			db.Init(config.MysqlDNS)
			// 3.注册服务
			err := pb.RegisterUserServiceHandler(service.Server(),
				handler.NewUserHandler(),
				server.InternalHandler(true))
			if err != nil {
				fmt.Println(err)
			}
		}),
		// 定义服务停止后的事情
		micro.AfterStop(func() error {
			logger.Info("user-srv服务停止打印...")
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Info("user-srv服务启动打印...")
			return nil
		}),
	)
	logger.Info("启动user-srv服务...")
	// 4.启动service
	if err := service.Run(); err != nil {
		logger.Panic("user-srv服务启动失败")
	}
}
