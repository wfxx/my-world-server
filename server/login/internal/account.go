package internal

import (
	"fmt"
	"my-world/server/mysql"
)

type Account struct {
	PlayerID uint   `gorm:"primary_key"`
	Username string `gorm:"not null;unique"`
	Password string
}

func checkUsername(username string) bool {
	if len(username) < 2 || len(username) > 12 {
		return false
	}
	return true
}

func checkPassword(password string) bool {
	if len(password) < 2 || len(password) > 12 {
		return false
	}
	return true
}

func newAccount(username string, password string) *Account {
	db := mysql.MysqlDB()
	var account = Account{Username: username, Password: password}
	err := db.Create(&account).Error
	if nil != err {
		return nil
	}
	return &account
}

func getAccount(Username string) *Account {
	var account Account
	db := mysql.MysqlDB()
	err := db.Where("Account = ?", Username).Limit(1).Find(&account).Error
	if nil != err {
		fmt.Println(err)
		return nil
	}
	fmt.Println("password:", account.Password)
	return &account
}
