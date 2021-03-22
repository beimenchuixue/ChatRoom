package message

type Kind uint

// 定义消息的类型
const (
	Login Kind = iota
	Response
	Register
	Exit
	UserUpLine
)

// 客户端与服务端传输消息的总协议
type Msg struct {
	Type Kind
	Data string
}

// 用户登录发送给服务器的认证信息（用户id和用户密码）
type LoginMsg struct {
	UserId int
	UserPwd string
}

// 响应用户登录状态的消息
type ResponseMsg struct {
	// 响应的状态码
	Code int
	// 响应的错误信息
	Error string
	// 在线用户列表
	OnlineUser []int
}

// 注册消息体
type RegisterMsg struct {
	UserId int
	UserPwd string
}

// UserStatus 用户状态类型
type UserStatus uint8

const (
	// 上线
	UpLine UserStatus = iota
	// 下线
	OffLine
)

// OnlineNotify 用户上线通知的消息体
type OnlineNotify struct {
	// 用户id
	UserId int
	// 用户状态
	UserStatus UserStatus
}



