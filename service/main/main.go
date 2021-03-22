package main

import (
	"chatroom/service/app"
)

func main() {
	// 运行服务
	chatApp := app.NewChatApp("tcp", "0.0.0.0:8888")
	chatApp.Run()
}
