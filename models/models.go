package models

import (
	"log"
	"strconv"
	"fmt"
	"github.com/JiangInk/market_monitor/setting"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(
		setting.DBSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DBSetting.User,
			setting.DBSetting.Password,
			setting.DBSetting.Host + strconv.Itoa(setting.DBSetting.Port),
			setting.DBSetting.DBName,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DBSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}