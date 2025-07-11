package task1

import (
	"gorm.io/gorm"
)

// 用户表
type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:20;not null"`
	Posts []Post
}

// 文章表
type Post struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserID   uint
	Title    string `gorm:"size:100;not null"`
	Content  string `gorm:"size:10000"`
	Comments []Comment
}

// 评论表
type Comment struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	PostID uint
	Comm   string `gorm:"size:1000;not null"`
}

func initTables(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	users := []User{
		{
			Name: "张三",
			Posts: []Post{
				{
					Title:   "张三的文章1",
					Content: "文章内容1",
					Comments: []Comment{
						{Comm: "评论1"},
						{Comm: "评论2"},
					},
				},
				{
					Title:   "张三的文章2",
					Content: "文章内容2",
					Comments: []Comment{
						{Comm: "评论1"},
						{Comm: "评论2"},
					},
				},
			},
		},
		{
			Name: "李四",
			Posts: []Post{
				{
					Title:   "李四的文章1",
					Content: "文章内容1",
					Comments: []Comment{
						{Comm: "评论1"},
						{Comm: "评论2"},
					},
				},
				{
					Title:   "李四的文章2",
					Content: "文章内容2",
					Comments: []Comment{
						{Comm: "评论1"},
						{Comm: "评论2"},
					},
				},
			},
		},
		{
			Name: "王五",
			Posts: []Post{
				{
					Title:   "王五的文章1",
					Content: "文章内容1",
					Comments: []Comment{
						{Comm: "评论1"},
						{Comm: "评论2"},
					},
				},
				{
					Title:   "王五的文章2",
					Content: "文章内容2",
					Comments: []Comment{
						{Comm: "评论1"},
						{Comm: "评论2"},
					},
				},
			},
		},
	}

	db.Create(&users)
}

func Run(db *gorm.DB) {
	initTables(db)
}
