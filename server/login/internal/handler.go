package internal

import (
	"crypto/md5"
	"encoding/hex"
	"my-world/server/hall"
	"my-world/server/msg"
	"reflect"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.RegisterRequest{}, handleRegister)
	handleMsg(&msg.LoginRequest{}, handleLogin)
}

func handleRegister(args []interface{}) {
	m := args[0].(*msg.RegisterRequest)
	a := args[1].(gate.Agent)

	// 数据库是否有该用户
	user := getAccount(m.Username)
	if user != nil {
		a.WriteMsg(&msg.RegisterResponse{Issuccess: false, Tips: "用户名已存在"})
		return
	}

	// 验证用户名是否符合规范
	if checkUsername(m.Username) {
		a.WriteMsg(&msg.RegisterResponse{Issuccess: false, Tips: "用户名不符合规范"})
		return
	}

	// 验证密码是否符合规范
	if checkPassword(m.Password) {
		a.WriteMsg(&msg.RegisterResponse{Issuccess: false, Tips: "密码不符合规范"})
		return
	}

	// 创建用户数据
	data := []byte(m.Password)
	var hash = md5.Sum(data)
	password := hex.EncodeToString(hash[:])
	newAccount(m.Username, password)
	hall.ChanRPC.Go("CreateUser")
	// a.WriteMsg(&msg.RegisterResponse{Issuccess: true, Tips: "用户注册成功"})
}

func handleLogin(args []interface{}) {
	m := args[0].(*msg.LoginRequest)
	a := args[1].(gate.Agent)

	// 数据库是否有该用户
	user := getAccount(m.Username)
	if user == nil {
		a.WriteMsg(&msg.LoginResponse{Issuccess: false, Tips: "该用户不存在"})
		return
	}
	// 校验密码是否正确
	data := []byte(m.Password)
	var hash = md5.Sum(data)
	password := hex.EncodeToString(hash[:])
	if password != user.Password {
		a.WriteMsg(&msg.LoginResponse{Issuccess: false, Tips: "用户名或密码错误"})
		return
	}
	a.WriteMsg(&msg.RegisterResponse{Issuccess: true, Tips: "用户登录成功"})
}
