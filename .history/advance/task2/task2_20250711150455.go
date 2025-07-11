package task2

import (
	"fmt"
	"go-gorm/advance/task1"

	"gorm.io/gorm"
)

func findAllPostAndCommentByName(db *gorm.DB, userName string) {
	users := []task1.User{}
	db.Debug().Model(&users).Preload("Posts").Where("name=?", userName).Find(&users)
	fmt.Printf("李四的信息:%#v\n", users)
}

func Run(db *gorm.DB) {
	findAllPostAndCommentByName(db, "李四")
}
