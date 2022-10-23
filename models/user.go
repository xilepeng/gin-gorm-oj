package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(100);" json:"name" `       // 用户名
	Password string `gorm:"column:password;type:varchar(32);" json:"password"` // 密码
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`       // 手机
	Mail     string `gorm:column:mail;type:varchar(100);" json:"mail"`         // 邮箱
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 用户的唯一标识
}

func (table *User) TableName() string {
	return "user"
}
