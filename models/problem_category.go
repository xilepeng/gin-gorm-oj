package models

import "gorm.io/gorm"

type ProblemCategory struct {
	ID            uint           `gorm:"primarykey;" json:"id"`
	CreatedAt     MyTime         `json:"created_at"`
	UpdatedAt     MyTime         `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	ProblemId     uint           `gorm:"column:problem_id;type:varchar(36);" json:"problem_id"`   // 问题的id
	CategoryId    uint           `gorm:"column:category_id;type:varchar(36);" json:"category_id"` // 分类的id
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"`                    // 关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
