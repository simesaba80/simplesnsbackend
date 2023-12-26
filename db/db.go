package db

import (
	"log"

	"snsback/utils/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	//todo 接続文字列を.envファイルにうつす
	dsn := config.MysqlTestUser + ":" + config.MysqlPassword + config.Mysqlconfig
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Post{})
	DB.AutoMigrate(&Friend{})
}
