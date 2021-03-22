package model

// User 用户对象，包含id 密码 名字三个字段
type User struct {
	UserId int	`json:"userId"`
	UserPwd string	`json:"userPwd"`
	UserName string	`json:"userName"`
}
