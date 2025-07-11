package task2

import (
	"fmt"
	"go-gorm/advance/task1"

	"gorm.io/gorm"
)

func findAllPostAndCommentByName(db *gorm.DB, userName string) {
	users := []task1.User{}
	db.Preload("Posts.Comments").Where("name=?", userName).Find(&users)
	fmt.Printf("%s的信息:%#v\n", userName, users)
}

func findMostCommentPost(db *gorm.DB) {}

func Run(db *gorm.DB) {
	findAllPostAndCommentByName(db, "王五")
}
