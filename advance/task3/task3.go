package task3

import (
	"fmt"
	"strconv"

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
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Title     string `gorm:"size:100;not null"`
	Content   string `gorm:"size:10000"`
	IsComment bool   `gorm:"column:is_comment;default:false;not null"`
	Comments  []Comment
}

// 评论表
type Comment struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	PostID uint
	Comm   string `gorm:"size:1000;not null"`
}

var POST_COUNT = 0

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	POST_COUNT++
	return nil
}

func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	var postID int64 = -1
	// 根据删除的Comment主键id查询Post主键
	tx.Model(&Comment{}).Select("post_id").Where("id=?", c.ID).Scan(&postID)
	var ct int64 = 0
	tx.Model(&Comment{}).Where("post_id=?", postID).Count(&ct)
	fmt.Println("count is :", ct)
	if ct == 1 {
		tx.Model(&Post{}).Where("id=?", postID).Update("is_comment", false)
		fmt.Println("change post status")
	}
	return nil
}

func CreatePost(db *gorm.DB) {
	ps := []Post{}
	for i := range 3 {
		p := Post{
			UserID:  uint(i + 1),
			Title:   "文章" + strconv.Itoa(i+1),
			Content: "内容" + strconv.Itoa(i+1),
		}
		ps = append(ps, p)
	}
	db.Create(&ps)

	fmt.Printf("created %d post\n", POST_COUNT)
}

func DeletePost(db *gorm.DB) {
	c := &Comment{ID: 2}
	db.Model(&Comment{}).Delete(&c)
}
