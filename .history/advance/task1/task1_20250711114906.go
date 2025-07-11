package task1

// 用户
type User struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:20;not null"`
}

type Post struct{}

type Comment struct{}
