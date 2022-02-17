package main

import (
	"demo1/models"
	"demo1/pkg/setting"
	"demo1/routers"
	"demo1/service/student_service"
	"fmt"
)

func main() {
	// 加载配置文件
	if err := setting.Init("conf/config.ini"); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 初始化数据库连接
	models.Setup()
	//程序关闭时关闭数据库连接
	defer models.Close()

	// 注册路由
	r := routers.SetupRouter()

	// 启动一个协程运行consumer方法
	go student_service.ChannelConsumer()

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
