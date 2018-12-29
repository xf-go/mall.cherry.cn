package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1)/cherry_mall?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "server:tijmaGudenumobae369uwginlepubjog@tcp(rm-2zen73w1368g0x25u7o.mysql.rds.aliyuncs.com)/cherry_user?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.SingularTable(true)

	DB = db
}
