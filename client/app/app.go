package app

import (
	"chatroom/client/view"
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

// 客户端运行
func (ca *ChatApp) Run() {
	router := view.Router{}
	router.IndexRouter()
}
