package task1

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:20;not null"`
}

// 文章表
type Post struct {
	ID int `gorm:"primaryKey;autoIncrement"`
}

// 评论表
type Comment struct{}
