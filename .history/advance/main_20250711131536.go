package main

import (
	"go-gorm/advance/task1"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dns := "root:Lq5942182!@tcp(192.168.1.168:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}

	task1.Run(db)
}
