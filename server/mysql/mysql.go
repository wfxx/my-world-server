package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
)

var (
	db *gorm.DB
)

func OpenDB() {
	fmt.Println("mysqldb->open db")
	username := "root"
	password := "xia12345"
	host := "127.0.0.1"
	port := 3306
	dbname := "myworld"
	// timeout := 10
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
	db1, err := gorm.Open("mysql", dsn)
	// db1, err := gorm.Open("mysql", "mike:123456@tcp(localhost:3306)/poker?parseTime=true")
	if err != nil {
		panic("connect db error")
	}
	db = db1
}

func MysqlDB() *gorm.DB {
	return db
}
