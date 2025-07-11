package task1

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model
	Name   string `gorm:"size:20;not null"`
	PostID int    // 文章外键
	Post
}

// 文章表
type Post struct {
	gorm.Model
	CommentID int //
	Comment
}

// 评论表
type Comment struct {
	gorm.Migrator
}
