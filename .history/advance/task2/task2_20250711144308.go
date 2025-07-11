package task2

import (
	"gorm.io/gorm"
)

func findAllPostAndCommentByName(db *gorm.DB, userName string) {
	user := []User
	// TODO: implement function logic
	db.Preload("Comments").Preload("Posts").Where("name=?", userName).Find(&user)
}

func Run(db *gorm.DB) {
	findAllPostAndCommentByName(db, "张三")
}
