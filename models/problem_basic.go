package models

import (
	"gorm.io/gorm"
)

// 1. gorm 字段忘记;结束
// 2. 数据库里命名和 models 里命名不一致
// Error 1054: Unknown column 'problem.category_id' in 'field list'

type ProblemBasic struct {
	gorm.Model

	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 问题表的唯一标识

	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"` // 关联问题分类表
	// CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"` // 分类id,已逗号分隔
	Title      string      `gorm:"column:title;type:char(255);" json:"title"`           // 文章标题
	Content    string      `gorm:"column:content;type:text;" json:"content"`            // 文章正文
	MaxRuntime int         `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"` // 最大运行时间
	MaxMem     int         `gorm:"column:max_mem;type:int(11);" json:"max_mem"`         // 最大运行内存
	TestCase   []*TestCase `gorm:"foreignKey:problem_identity;references:identity"`     // 关联测试用例表
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	// data := make([]*Problem, 0)
	// DB.Find(&data)

	// for _, v := range data {
	// 	fmt.Printf("Problem ==> %v \n", v)
	// }
	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity) // 错误：cb ——> cd
	}
	// if categoryIdentity != "" {
	// 	tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
	// 		Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
	// }
	return tx
}
