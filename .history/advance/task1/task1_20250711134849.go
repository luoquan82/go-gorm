package task1

import "gorm.io/gorm"

// 用户表
type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:20;not null"`
	Posts []Post `gorm:"foreignKey:UserID"`
}

// 文章表
type Post struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserID   uint
	Title    string    `gorm:"size:100;not null"`
	Content  string    `gorm:"size:10000"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

// 评论表
type Comment struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Comm   string `gorm:"size:1000;not null"`
	PostID uint
}

func initTables(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	articles := []User{
		{
			Name: "张三",
			Posts: []Post{
				{
					Title:   "张三的文章1",
					Content: "张三文章1的内容",
					Comments: []Comment{
						{Comm: "评论1"},
						{Comm: "评论2"},
						{Comm: "评论3"},
						{Comm: "评论4"},
					},
				},
			},
		},
	}
}

func Run(db *gorm.DB) {
	initTables(db)
}
