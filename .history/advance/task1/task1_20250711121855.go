package task1

// 用户表
type User struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:20;not null"`
}

// 文章表
type Post struct {
	ID int `gorm:"primaryKey;autoIncrement"`
}

// 评论表
type Comment struct{}
