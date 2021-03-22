package dao

import (
	"chatroom/service/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var (
	ctx = context.Background()
	// UserTableName 用户表名
	UserTableName = "users"
)

// UserDao 处理用户登录、注册逻辑
type UserDao struct {
	// redis连接对象
	Cli *redis.Client
}

func NewUserDao()  *UserDao {
	return &UserDao{Cli: redisCli}
}

// 查询用户
func (u *UserDao) getUserById(userId int) (user *model.User, errCode int, err error) {
	// 1. 查询结果
	res, err := u.Cli.HGet(ctx, UserTableName, strconv.Itoa(userId)).Result()
	if err != nil {
		// 2. 判断查询结果是否是空结果
		if err == redis.Nil {
			return nil, UserNotExitErr, UserError[UserNotExitErr]
		}
		return nil, NotErr, err
	}

	// 3. 对结果进行反序列化成需要的User数据结构
	user = &model.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		return nil, NotErr, err
	}
	return
}

// Login 用户登录逻辑
func (u *UserDao) Login(userId int, userPwd string)	(user *model.User, errCode int, err error) {
	// 1. 查询用户
	user, errCode, err = u.getUserById(userId)
	if err != nil {
		return nil, errCode, err
	}
	// 验证用户密码，成功则返回用户实例，失败则返回用户密码错误信息
	if user.UserPwd == userPwd {
		return user, NotErr, nil
	} else {
		return nil, UserPwdErr, UserError[UserPwdErr]
	}
}

// AddUser 增加用户
func (u *UserDao) AddUser(userId int, userPwd string) (errCode int, err error)  {
	// 1. 查询一下用户是否已经存在
	_, errCode, err = u.getUserById(userId)
	if err != nil {
		if errCode == UserNotExitErr {
			// 1. 构建用户信息数据结构
			user := model.User{
				UserId: userId,
				UserPwd: userPwd,
			}
			// 2. 序列化为字符串
			userJson, err := json.Marshal(&user)
			if err != nil {
				return 0, err
			}
			// 3. 存入redis
			_, err = u.Cli.HSet(ctx, UserTableName, strconv.Itoa(userId), userJson).Result()

			if err != nil {
				return 0, err
			}
			return 0, err
		}
	}
	return
}