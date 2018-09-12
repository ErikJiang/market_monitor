package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JiangInk/market_monitor/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(
		config.DBSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DBSetting.User,
			config.DBSetting.Password,
			config.DBSetting.Host+":"+strconv.Itoa(config.DBSetting.Port),
			config.DBSetting.DBName,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.DBSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
