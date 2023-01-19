package internal

import (
	"my-world/server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("CreateUser", rpcCreateUser)
}

func rpcCreateUser(args []interface{}) {
	agent := args[0].(gate.Agent)

	log.Debug("随机生成角色")

	m := &msg.RegisterResponse{Issuccess: true, Tips: "用户注册成功"}
	agent.WriteMsg(m)
}
