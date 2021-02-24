package bootstrap

import (
	"fmt"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	."go-test/config"
)

var Db *gorm.DB

func InitDB() {
	var err error
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.GetString("database.mysql.user"),
		Config.GetString("database.mysql.password"),
		Config.GetString("database.mysql.host"),
		Config.GetString("database.mysql.port"),
		Config.GetString("database.mysql.database")))

	if err != nil {
		panic("failed to connect database")
	}
	if Config.GetString("app.env") != "release" {
		Db.LogMode(true)
	}
	Db.SingularTable(true)
}
