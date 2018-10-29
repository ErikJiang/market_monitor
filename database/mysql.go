package database

import (
	"fmt"
	"strconv"

	"github.com/JiangInk/market_monitor/config"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Setup() {
	var err error
	DB, err = gorm.Open(
		config.DBSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DBSetting.User,
			config.DBSetting.Password,
			config.DBSetting.Host+":"+strconv.Itoa(config.DBSetting.Port),
			config.DBSetting.DBName,
		),
	)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.DBSetting.TablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
