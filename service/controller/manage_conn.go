package controller

import (
	"chatroom/common/message"
	"chatroom/service/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

// UserConn 用户单个连接信息
type UserConn struct {
	Conn net.Conn
	UserId int
	UserStatus message.UserStatus
}

// NotifyMe 将对应用户的id通知到我自己
func (uc *UserConn) NotifyMe(userId int) error {
	// 1. 构建用户上线消息体
	var notifyMsg message.OnlineNotify
	notifyMsg.UserStatus = message.UpLine
	notifyMsg.UserId = userId
	// 2. 构建消息
	var msg message.Msg
	msg.Type = message.UserUpLine
	notifyData, err := json.Marshal(&notifyMsg)
	if err != nil {
		return err
	}
	msg.Data = string(notifyData)
	// 3. 发送消息
	tf := utils.NewTransfer(uc.Conn)
	err = tf.WriteData(&msg)
	if err != nil {
		return err
	}
	return nil
}

var (
	// userAllConn 存放所有连接到服务器的连接，是一个单独对象
	UserAllConn *userAllConn
)

// UserAllConn 保存用户所有连接
type userAllConn struct {
	allConn map[int]*UserConn
}

func init() {
	// 初始化全局的唯一保存所有连接到额对象
	UserAllConn = &userAllConn{
		allConn: make(map[int]*UserConn, 20),
	}
}

// List 列出当前所有连接
func (uac *userAllConn) List() map[int]*UserConn {
	return uac.allConn
}

// GetConnById 通过用户id获取单个连接对象
func (uac *userAllConn) GetConnById(userId int)(userConn *UserConn, err error)  {
	conn, ok := uac.allConn[userId]
	if !ok {
		err = errors.New(fmt.Sprintf("当前用户id%v对应的连接不存在", userId))
		return
	}
	return conn, nil
}

// AddOrUpdate 添加或者更新连接
func (uac *userAllConn)AddOrUpdate(userConn *UserConn) {
	uac.allConn[userConn.UserId] = userConn
}

// NotifyAllUser 将用户登录状态发送给所有用户
func (uac *userAllConn) NotifyAllUser(userId int)  {
	// 1. 遍历所有连接，忽略自身连接，将自己上线消息发送给所有的在线连接
	for id, userConn := range uac.allConn {
		// 排除给自己发送连接
		if id == userId {
			continue
		}
		err := userConn.NotifyMe(userId)
		if err != nil {
			fmt.Println(err)
		}
	}
}
