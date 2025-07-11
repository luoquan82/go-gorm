package task1

import "gorm.io/gorm"

// 用户表
type User struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"size:20;not null"`
	PostID int    // 文章外键
	Post
}

// 文章表
type Post struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"size:100;not null"`
	Content   string `gorm:"size:10000"`
	CommentID int    // 评论外键
	Comment
}

// 评论表
type Comment struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Comment string `gorm:"size:1000;not null"`
}

func initTables(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func Run(db *gorm.DB) {
	initTables(db)
}
