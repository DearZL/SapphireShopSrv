package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uint64 `gorm:"primaryKey"`
	CreatedAt int64  `gorm:"column:create_time;autoCreateTime;not null"`
	UpdatedAt int64  `gorm:"column:update_time;autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt
}

type User struct {
	BaseModel
	UserId   string `gorm:"index:idx_user_id;unique;not null"`
	Email    string `gorm:"index:idx_email;unique;type:varchar(30);not null"`
	Password string `gorm:"type:varchar(30);not null"`
	UserName string `gorm:"type:varchar(20);not null"`
	Birthday int64  `gorm:"not null"`
	Sex      uint8  `gorm:"column:sex;default:1;type:varchar(6) comment '1表示男,2表示女';not null"`
	Role     uint8  `gorm:"column:role;default:1;type:int comment '1表示普通用户, 2表示管理员';not null"`
}
