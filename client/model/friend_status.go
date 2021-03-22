package model

// StatusType 用户在线状态类型
type StatusType uint8

const (
	// 上线
	UpLine StatusType = iota
	// 下线
	OffLine
)

// User 用户状态信息，包含用户id、名称、状态
type User struct {
	Code int
	Name string
	Status StatusType
}
