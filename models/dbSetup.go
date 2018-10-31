package models

import (
	"fmt"
	"strconv"

	"github.com/JiangInk/market_monitor/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

// DB 当前数据库连接
var DB *gorm.DB

// Setup MySQL 数据库配置
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
	// migrate 迁移
	DB.Set("grom:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin").AutoMigrate(&User{})
	DB.Model(&User{}).AddUniqueIndex("uk_email", "email")
}
