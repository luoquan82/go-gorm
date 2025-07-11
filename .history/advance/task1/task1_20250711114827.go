package task1

// 用户
type User struct {
	ID int `gorm:"primaryKey;autoIncrement"`
}

type Post struct{}

type Comment struct{}
