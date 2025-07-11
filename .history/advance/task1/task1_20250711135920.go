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

	users := []User{
		{Name: "张三"},
		{Name: "李四"},
		{Name: "王五"},
		{Name: "赵六"},
		{Name: "郑七"},
	}

	users[1].Posts = []Post{
		{Title: "张三的文章1"},
		{Title: "张三的文章2"},
		{Title: "张三的文章3"},
	}

	db.Create(&users)
}

func Run(db *gorm.DB) {
	initTables(db)
}
