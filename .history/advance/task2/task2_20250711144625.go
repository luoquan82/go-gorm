package task2

import (
	"fmt"
	"go-gorm/advance/task1"

	"gorm.io/gorm"
)

func findAllPostAndCommentByName(db *gorm.DB, userName string) {
	users := []task1.User{}
	// TODO: implement function logic
	db.Preload("Posts").Where("name=?", userName).Find(&users)
	fmt.Println("张三的信息:%#v", users)
}

func Run(db *gorm.DB) {
	findAllPostAndCommentByName(db, "张三")
}
