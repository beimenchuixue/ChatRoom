package main

import "chatroom/client/app"

func main() {
	// 运行客户端
	ChatAppCli := app.NewChatApp("tcp", "127.0.0.1:8888")
	ChatAppCli.Run()
}
