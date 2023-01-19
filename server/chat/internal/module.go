package internal

import (
	"my-world/server/base"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	InitGameTables()
	InitRooms()

}

func InitGameTables() {
	//db:=mysql.MysqlDB()
	//db.AutoMigrate(&PlayerBaseInfo{})
}

func (m *Module) OnDestroy() {

}
