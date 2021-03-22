package view

import (
	"log"
)

// index页用户不同的选择
const (
	LoginChoice  = iota + 1
	RegisterChoice
	ExitChoice
)

// Home页用户不同的选择
const (
	OnlineUserChoice  = iota + 1
	SendMsgChoice
	MsgListChoice
	ExitHomeChoice
)

type Router struct {
}

// MenuRouter 根据主菜单的选择，映射到不同的处理流程
func (r *Router)IndexRouter()  {
	// 1. 调用用户视图，返回用户的选择
	userV := UserView{}
	// 2. 获取用户的选择
	indexLabel:
	choiceKey := userV.IndexView()
	switch choiceKey {
	// 登录逻辑
	case LoginChoice:
		// ok判断用户是否验证成功
		ok, err := userV.LoginView()
		if err != nil {
			log.Println(err)
			return
		}
		if ok {
			r.HomeRouter()
		}

	// 注册逻辑
	case RegisterChoice:
		_, err := userV.RegisterView()
		if err != nil {
			log.Println(err)
			return
		}
		goto indexLabel
	// 退出逻辑
	case ExitChoice:
		userV.ExistView()
	}
}

func (r Router) HomeRouter() {
	// 调用用户视图，返回用户的选择
	userV := UserView{}
	// 获取用户的选择
	choiceKey := userV.HomeView()
	switch choiceKey {

	// 显示在线用户
	case OnlineUserChoice:
		err := userV.OnlineUserListView()
		if err != nil {
			log.Println(err)
			return
		}
	// 发送消息
	case SendMsgChoice:
		err := userV.SendMsgView()
		if err != nil {
			log.Println(err)
			return
		}
	// 信息列表
	case MsgListChoice:
		err := userV.MsgListView()
		if err != nil {
			log.Println(err)
			return
		}
	// 退出home页面逻辑
	case ExitHomeChoice:
		userV.ExitHomeView()
	}
}