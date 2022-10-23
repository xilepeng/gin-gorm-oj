package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	ID         uint   `gorm:"primarykey;" json:"id"`
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`       // 问题表的唯一标识
	CategoryId string `gorm:"column:category_id;type:varchar(255)" json:"category_id"` // 分类id,已逗号分隔
	Title      string `gorm:"column:title;type:char(255)" json:"title"`                // 文章标题
	Content    string `gorm:"column:content;type:text" json:"content"`                 // 文章正文
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11)" json:"max_runtime"`      // 最大运行时间
	MaxMem     int    `gorm:column"max_mem;type:int(11)" json:"max_mem"`               // 最大运行内存
}

func (table *Problem) TableName() string {
	return "problem"
}

func GetProblemList() {

	data := make([]*Problem, 0)
	DB.Find(&data)

	for _, v := range data {
		fmt.Printf("Problem ==> %v \n", v)
	}
}
