package task1

import "gorm.io/gorm"

// 用户表
type User struct {
	ID     int    `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"size:20;not null"`
	PostID int    // 文章外键
	Post
}

// 文章表
type Post struct {
	gorm.Model
	// ID        int `gorm:"primaryKey;autoIncrement"`
	CommentID int //
	Comment
}

// 评论表
type Comment struct {
	ID int `gorm:"primaryKey;autoIncrement"`
}
