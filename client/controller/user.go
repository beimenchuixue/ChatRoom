package controller

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"errors"
	"fmt"
)

// 处理用户信息

type User struct {
	UserId int
	UserPwd string
}

func (u *User)Login() (ok bool, err error) {
	// 向服务器发送用户身份认证信息 用户id 和用户密码

	// 1. 构建发送的消息体
	var loginMsg message.LoginMsg
	loginMsg.UserId = u.UserId
	loginMsg.UserPwd = u.UserPwd

	// 2.发送消息
	tf := utils.NewTransfer()
	err = tf.SendData(u, message.Login)
	if err != nil {
		return
	}

	// 2. 接收消息
	data, dataType, err := tf.RecvData()
	if err != nil {
		return
	}

	// 3. 验证响应结果
	if dataType == message.Response {
		var res message.ResponseMsg
		err = json.Unmarshal([]byte(data), &res)
		if err != nil {
			return
		}
		if res.Code == 0 {
			for _, v := range res.OnlineUser {
				fmt.Printf("%v 在线\n", v)
			}
			ok = true
		} else {
			err = errors.New(res.Error)
		}
	} else {
		err = errors.New(fmt.Sprintf("login接收消息类型为%v, 但接收%v", message.Response, dataType) )
		return
	}
	return
}

// Register 用户注册
func (u *User) Register() (ok bool, err error){
	// 1. 构建注册消息体
	var registerMsg message.RegisterMsg
	registerMsg.UserId = u.UserId
	registerMsg.UserPwd = u.UserPwd
	// 2.发送消息
	tf := utils.NewTransfer()
	err = tf.SendData(u, message.Register)
	if err != nil {
		return
	}

	// 2. 接收消息
	data, dataType, err := tf.RecvData()
	if err != nil {
		return
	}
	// 3. 解析注册响应结果
	if dataType == message.Register {
		var res message.ResponseMsg
		err = json.Unmarshal([]byte(data), &res)
		if err != nil {
			return
		}
		if res.Code == 0 {
			fmt.Println("注册成功，请登录")
			return true, nil
		}  else {
			return false, errors.New(res.Error)
		}
	} else {
		err = errors.New( fmt.Sprintf("register接收消息类型为%v, 但接收%v", message.Register, dataType) )
		return
	}
}