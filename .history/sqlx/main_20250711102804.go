package main

import "github.com/jmoiron/sqlx"

var db *sqlx.DB
var err error

func init() {
	dns := "root:Lq5942182!@tcp(192.168.1.168:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
}

func main() {}
