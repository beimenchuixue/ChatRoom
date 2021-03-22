package controller

import (
	"chatroom/common/message"
	"chatroom/service/dao"
	"encoding/json"
	"net"
)

// User 处理用户业务
type User struct {
	Conn net.Conn
}

// Login 处理用户登录逻辑
func (u *User)LoginController(msgData *message.Msg) (msg *message.Msg, err error) {
	var loginMsg message.LoginMsg
	//1. 获取用户登录的用户名和密码
	err = json.Unmarshal([]byte(msgData.Data), &loginMsg)
	if err != nil {
		return nil, err
	}

	var response message.ResponseMsg
	// 1.用户身份认证
	userDao := dao.NewUserDao()
	_, errCode, err := userDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		// 登录失败
		response.Code = errCode
		if errCode != dao.NotErr {
			response.Error = dao.UserError[errCode].Error()
		}
	} else {
		// 构建一个连接，将登录成功的连接放入总连接池
		userC := &UserConn{
			Conn:   u.Conn,
			UserId: loginMsg.UserId,
			UserStatus: message.UpLine,
		}
		UserAllConn.AddOrUpdate(userC)
		// 2. 构建响应信息
		for v, _ := range UserAllConn.List() {
			response.OnlineUser = append(response.OnlineUser, v)
		}
		// 3. 通知所有在线连接
		UserAllConn.NotifyAllUser(loginMsg.UserId)
	}
	resMsg, err := json.Marshal(&response)
	if err != nil {
		return nil, err
	}
	// 3. 填充消息体
	msgData.Type = message.Response
	msgData.Data = string(resMsg)
	// 4. 发送消息体
	return msgData, err
}

// RegisterController 用户注册
func (u *User) RegisterController(msgData *message.Msg) (msg *message.Msg, err error) {
	var registerMsg message.RegisterMsg
	//1. 获取注册的用户id和密码
	err = json.Unmarshal([]byte(msgData.Data), &registerMsg)
	if err != nil {
		return nil, err
	}

	userDao := dao.NewUserDao()
	errCode, err := userDao.AddUser(registerMsg.UserId, registerMsg.UserPwd)
	var response message.ResponseMsg
	if err != nil {
		response.Code = errCode
		response.Error = err.Error()
	}

	// 2. 构建响应信息
	resMsg, err := json.Marshal(&response)
	if err != nil {
		return nil, err
	}
	// 3. 填充消息体
	msgData.Type = message.Register
	msgData.Data = string(resMsg)
	// 4. 发送消息体
	return msgData, err
}