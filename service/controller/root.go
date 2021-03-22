package controller

import (
	"chatroom/common/message"
	"chatroom/service/utils"
	"fmt"
	"log"
	"net"
)

// 总控制器处理每一个客户端连接
func RootController(conn net.Conn){
	defer conn.Close()
	loop := true
	tf := utils.NewTransfer(conn)
	// 循环处理客户端的消息
	for loop {
		// 1. 读取消息
		readMsg, err := tf.ReadData()
		// 输出接收到的信息
		log.Println(readMsg)
		if err != nil {
			log.Println(err)
			return
		}

		// 2. 处理消息，根据消息的类型不同，路由到不同的控制器
		var sendMsg *message.Msg
		// 依据不同消息类型，交给不同的控制器
		fmt.Println(readMsg.Type)
		switch readMsg.Type {
		// 处理登录逻辑
		case message.Login:
			userControl := User{Conn: conn}
			sendMsg, err = userControl.LoginController(readMsg)
			if err != nil {
				log.Println(err)
				return
			}

			// 处理注册
		case message.Register:
			userControl := User{Conn: conn}
			sendMsg, err = userControl.RegisterController(readMsg)
			if err != nil {
				log.Println(err)
				return
			}

			// 处理退出逻辑
		case message.Exit:
			loop = false
		default:
			fmt.Println("TODO 等待实现")
		}
		// 输出即将发送的信息
		log.Println(sendMsg)
		// 3. 发送消息
		err = tf.WriteData(sendMsg)
		if err != nil {
			log.Println(err)
			return
		}
	}

}
