package dao

import "errors"

// 自定义跟用户业务逻辑相关的错误
var (
	UserError map[int]error
)

const (
	NotErr = iota
	UserNotExitErr
	UserPwdErr
	UserExit
)

func init() {
	UserError = make(map[int]error, 20)
	// 用户登录时候错误逻辑
	UserError[UserNotExitErr] = errors.New("用户不存在")
	UserError[UserPwdErr] = errors.New("用户密码错误")

	// 用户注册错误逻辑
	UserError[UserExit] = errors.New("用户存在")
}



