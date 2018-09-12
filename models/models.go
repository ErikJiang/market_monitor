package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	ID			int	`gorm:"primary_key" json:"id"`
	CreatedOn	int	`gorm:"created_on"`
	ModifiedOn	int `gorm:"modified_on"`
	DeletedOn	int `gorm:"deleted_on"`
}

func Setup() {
	var err error
	db, err = gorm.Open()
}