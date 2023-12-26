package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MysqlTestUser string
	MysqlPassword string
	Mysqlconfig   string
)

// .envを呼び出します。
func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("読み込み出来ませんでした: %v", err)
	}
	MysqlTestUser = os.Getenv("MYSQLTESTUSER")
	MysqlPassword = os.Getenv("MYSQLPASSWORD")
	Mysqlconfig = os.Getenv("MYSQLCONFIG")
}
