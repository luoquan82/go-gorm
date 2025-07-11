package task1

import "gorm.io/gorm"

// 用户表
type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:20;not null"`
	Posts []Post
}

// 文章表
type Post struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Title   string `gorm:"size:100;not null"`
	Content string `gorm:"size:10000"`
	UserID  uint
	User    User
}

// 评论表
type Comment struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Comment string `gorm:"size:1000;not null"`
	PostID  uint
	Post    Post
}

func initTables(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	articles:=[]User{
		{Name: "张三" }
	}
}

func Run(db *gorm.DB) {
	initTables(db)
}
