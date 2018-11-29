package models

import (
	"github.com/jinzhu/gorm"
)

// Task 任务表 model 定义
type Task struct {
	gorm.Model
	User    User    `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID  int     `gorm:"column:userId;not null"`
	Type    string  `sql:"type:ENUM('TICKER', 'OTHER')"`
	Status  string  `sql:"type:ENUM('ENABLE', 'DISABLE')"`
	Rules    string  `gorm:"column:rules;type:varchar(200);not null"`
}
