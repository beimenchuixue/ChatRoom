package view

import (
	"chatroom/client/controller"
	"fmt"
	"log"
	"os"
)

type UserView struct {
}

func (ur *UserView)LoginView() (ok bool, err error) {
	// 1. 终端获取用户输入的用户id和用户密码信息，用于验证用户身份
	var userId int
	var userPwd string
	tryCount := 3
	for tryCount >= 1 {

		fmt.Print("请输入用户ID:")
		fmt.Scanln(&userId)
		fmt.Print("请输入用户密码:")
		fmt.Scanln(&userPwd)
		// 2. 向服务端发送用户名和密码
		userC := controller.User{UserId: userId, UserPwd: userPwd}
		ok, err = userC.Login()
		if err != nil {
			fmt.Printf("%v, 剩余尝试机会%v次\n", err, tryCount-1)
		}
		if ok {
			break
		}
		tryCount --
	}
	return
}

// IndexMenu 主页获取用户的选择
func (ur *UserView)IndexView() (choiceKey int) {
	var loop = true
	fmt.Println("--------聊天系统---------")
	fmt.Println("		 1. 登录")
	fmt.Println("		 2. 注册")
	fmt.Println("		 3. 退出")
	fmt.Print("请选择(1-3): ")
	for loop {
		_, err := fmt.Scanln(&choiceKey)
		if err != nil {
			log.Println(err)
			fmt.Print("输入错误，请重新输入(1-3):")
			continue
		}
		// 验证用户是否选择 1 2 3
		if choiceKey == LoginChoice || choiceKey == RegisterChoice || choiceKey == ExitChoice {
			loop = false
		} else {
			fmt.Print("输入错误，请重新输入(1-3):")
		}
	}
	return
}

// RegisterView 用户注册页面
func (ur *UserView) RegisterView() (ok bool, err error)  {
	// 1. 获取终端用户输入
	var userId int
	var userPwd string
	fmt.Println("-----------用户注册-----------")
	fmt.Print("请输入用户ID:")
	fmt.Scanln(&userId)
	fmt.Print("请输入用户密码:")
	fmt.Scanln(&userPwd)
	// 2.向服务端发送用户名和密码，交给控制层
	userC := controller.User{UserId: userId, UserPwd: userPwd}
	ok, err = userC.Register()
	if err != nil {
		return false, err
	}
	return
}

// ExistView用户退出界面
func (ur *UserView) ExistView()  {
	fmt.Println("退出聊天系统")
	os.Exit(0)
}

// IndexMenu 主页获取用户的选择
func (ur *UserView)HomeView() (choiceKey int) {
	var loop = true
	fmt.Println("---用户XXX---聊天主页---------")
	fmt.Println("		 1. 显示在线用户列表")
	fmt.Println("		 2. 发送消息")
	fmt.Println("		 3. 消息列表")
	fmt.Println("		 4. 退出系统")
	fmt.Print("请选择(1-4): ")
	for loop {
		_, err := fmt.Scanln(&choiceKey)
		if err != nil {
			log.Println(err)
			fmt.Print("输入错误，请重新输入(1-4):")
			continue
		}
		// 验证用户是否选择 1 2 3 4
		if choiceKey == LoginChoice || choiceKey == RegisterChoice ||
			choiceKey == ExitChoice || choiceKey == 4 {
			loop = false
		} else {
			fmt.Print("输入错误，请重新输入(1-3):")
		}
	}
	return
}

// OnlineUserListView 显示在线用户列表
func (ur UserView) OnlineUserListView() (err error) {
	return
}

// SendMsgView 用户发送消息界面
func (ur UserView) SendMsgView() (err error) {
	return
}

// MsgListView 用户消息列表
func (ur UserView) MsgListView() (err error) {
	return
}


// ExitHomeView 用户退出家页面
func (ur UserView) ExitHomeView()  {
	return
}