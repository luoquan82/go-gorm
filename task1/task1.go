package task1

import "gorm.io/gorm"

type Student struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:100;not null"`
	Age   int    `gorm:"not null"`
	Grade string `gorm:"size:50;not null"`
}

func Run(db *gorm.DB) {
	// 创建Students表
	db.AutoMigrate(&Student{})
	s := &Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}

	// 创建学生张三 create与save区别。create仅做新增，不做更新(update)操作
	db.Create(s)

	// 查询 students 表中所有年龄大于 18 岁的学生信息
}
