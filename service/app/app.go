package app

import (
	"chatroom/service/controller"
	"log"
	"net"
)

// ChatApp 包含一个一个app的配置初始化和运行方法
type ChatApp struct {
	network, address string
}

// NewChatApp 创建 ChatApp实例的工厂函数
func NewChatApp(network, address string ) *ChatApp {
	return &ChatApp{
		network: network,
		address: address,
	}
}

func (ca *ChatApp)Run()  {
	// 1. 启动监听
	listen, err := net.Listen(ca.network, ca.address)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	// 2. 接受连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// 3. 将连接交给总控制器
		go controller.RootController(conn)
	}
}
