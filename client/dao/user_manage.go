package dao

import (
	"chatroom/client/model"
)

// Users 保存所有用户状态信息
type Users map[int]*model.User

// UpdateOrAddUserStatus 更新或添加用户状态
func (u *Users) UpdateOrAddUserStatus(user *model.User) {
	m, ok := (*u)[user.Code]
	if !ok {
		// 不存在则添加
		(*u)[user.Code] = user
	} else {
		// 存在则更新
		m.Status = user.Status
	}
}

// OnlineUser 列出所有在线用户
func (u *Users) OnlineUser() []*model.User {
	onlineUser := make([]*model.User, 0, 20)
	for _, user := range *u {
		if user.Status == model.UpLine {
			onlineUser = append(onlineUser, user)
		}
	}
	return onlineUser
}