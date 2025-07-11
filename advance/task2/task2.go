package task2

import (
	"fmt"
	"go-gorm/advance/task1"

	"gorm.io/gorm"
)

type Result struct {
	task1.Post
	num int
}

func FindAllPostAndCommentByName(db *gorm.DB, userName string) {
	users := []task1.User{}
	db.Preload("Posts.Comments").Where("name=?", userName).Find(&users)
	fmt.Printf("%s的信息:%#v\n", userName, users)
}

func FindMostCommentPost(db *gorm.DB) {
	result := Result{}
	db.Raw(`
	select p.*,count(*) as num
	from posts p left join
	comments c on p.id=c.post_id
	group by p.id
	order by num desc limit 1
	`).Scan(&result)
	fmt.Printf("Result: %#v\n", result)
}
