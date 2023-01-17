package gate

import (
	"my-world/server/game"
	"my-world/server/msg"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.TosChat{}, game.ChanRPC)
}
