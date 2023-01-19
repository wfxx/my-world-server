package internal

import (
	"my-world/server/msg"
	"time"

	"github.com/name5566/leaf/gate"
	g "github.com/name5566/leaf/go"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/util"
)

var (
	userlist = make(map[uint]*User)
)

const (
	offline = iota
	online
	ingame
)

// 坐标
type Position struct {
	pname string
	px    float64
	py    float64
}

// 用户
type User struct {
	gate.Agent
	*g.LinearContext
	state int

	userInfo *UserInfo
	userPos  *Position
}

func (user *User) login(userID uint) {

	userInfo := new(UserInfo)
	user.userInfo = userInfo

	skeleton.Go(func() {
		err := userInfo.initValue(userID)
		if err != nil {
			log.Error("init acc %v data error: %v", userID, err)
			userInfo = nil
			user.WriteMsg(&msg.LoginFaild{Code: msg.LoginFaild_InnerError})
			player.Close()
			return
		}
	}, func() {
		// network closed
		if player.state == userLogout {
			player.logout(playerID)
			return
		}

		// db error
		player.state = userGame
		if playerBaseInfo == nil {
			return
		}

		// ok
		player.playerBaseInfo = playerBaseInfo
		playerID2Player[playerID] = player
		//player.UserData().(*AgentInfo).userID = userData.UserID
		player.onLogin()
		player.autoSaveDB()
	})

}

func CreatePlayer(playerID uint) error {
	err := CreateUserInfo(playerID)
	return err
}

func (player *Player) isOffline() bool {
	return player.state == userLogout
}

func (player *Player) logout(playerID uint) {

}

func (player *Player) autoSaveDB() {
	const duration = 5 * time.Minute
	// save
	player.saveDBTimer = skeleton.AfterFunc(duration, func() {
		data := util.DeepClone(player.playerBaseInfo)
		player.Go(func() {
			err := data.(*PlayerBaseInfo).saveValue()
			if err != nil {
				log.Error("save user %v data error: %v", player.playerBaseInfo.PlayerID, err)
			}

		}, func() {
			player.autoSaveDB()
		})
	})
}

func (player *Player) onLogin() {

}

func (player *Player) onLogout() {

}
