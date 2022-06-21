package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB情報
const (
	USER   string = "root"
	PASS   string = "deeptrack"
	HOST   string = "mysql-container"
	DBNAME string = "DeepTrack"
	OPTION string = "charset=utf8&parseTime=True&loc=Local"
)

//データベース接続
func DatabaseConnection() *gorm.DB {
	connection := USER + ":" + PASS + "@tcp(" + HOST + ")/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
