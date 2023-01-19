package internal

import (
	"fmt"
	"my-world/server/mysql"
)

type UserInfo struct {
	UserID    uint   `gorm:"primary_key"`
	UserNick  string `gorm:"not null"`
	HeadUrl   string
	BirthDay  string
	TelePhone string
	Address   string
}

func (userInfo *UserInfo) initValue(userID uint) error {
	mysql := mysql.MysqlDB()

	err := mysql.Where("UserID = ?", userID).Limit(1).Find(&userInfo).Error
	if nil != err {
		fmt.Println(err)
		return fmt.Errorf("get UserInfo id error: %v", err)
	}

	return nil
}

func (userInfo *UserInfo) saveValue() error {
	mysql := mysql.MysqlDB()
	err := mysql.Save(&userInfo).Error

	if nil != err {
		fmt.Println(err)
		userInfo = nil
		return fmt.Errorf("get UserInfo id error: %v", err)
	}

	return nil
}

func CreateUserInfo(userID uint) error {
	userInfo := new(UserInfo)

	mysql := mysql.MysqlDB()
	err := mysql.Save(&userInfo).Error

	if nil != err {
		fmt.Println(err)
		return fmt.Errorf("get UserInfo id error: %v", err)
	}

	return nil
}
