package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:123@tcp(127.0.0.1:3306)/cahoot?charset=utf8mb4&parseTime=True"

func InitDb() {
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(db)
	}
}
