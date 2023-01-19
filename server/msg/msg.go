package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = protobuf.NewProcessor()
)

// 注册protobuf消息
func init() {
	Processor.SetByteOrder(true)
	// 登录服
	Processor.Register(&RegisterRequest{})
	Processor.Register(&RegisterResponse{})
	Processor.Register(&LoginRequest{})
	Processor.Register(&LoginResponse{})
}
