package main

import "github.com/zeromicro/go-zero/core/stores/sqlx"

func init() {
	dns := "root:Lq5942182!@tcp(192.168.1.168:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	sqlx.Open(sqlx.MySQL, dns)
}

func main() {

}
