package task1

import (
	"fmt"

	"gorm.io/gorm"
)

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
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	fmt.Println(students)

	// 将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
	db.Where("name= ?", "张三").Model(&Student{}).Update("grade", "四年级")

	// 删除 students 表中年龄小于 15 岁的学生记录
	db.Debug().Where("age < ?", 15).Delete(&Student{})
}
