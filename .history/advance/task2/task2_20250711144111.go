package task2

import (
	"gorm.io/gorm"
)

func findAllPostAndCommentByName(db *gorm.DB, userName string) {
	user := []User{}
	// TODO: implement function logic
	db.Preload("Comment").Preload("Post").Where("name=?", userName).Find()
}

func Run(db *gorm.DB) {
	findAllPostAndCommentByName(db, "张三")
}
