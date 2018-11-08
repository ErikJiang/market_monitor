package config

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

// server 服务基本配置结构
type server struct {
	RunMode   string `mapstructure:"runMode"`
	Port      int    `mapstructure:"port"`
	JWTSecret string `mapstructure:"jwtSecret"`
	JWTExpire int64  `mapstructure:"jwtExpire"`
}

// ServerConf 服务基本配置
var ServerConf = &server{}

// email 邮箱配置结构
type email struct {
	ServName         string `mapstructure:"servName"`
	UserName         string `mapstructure:"userName"`
	Password         string `mapstructure:"password"`
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	ContentTypeHTML  string `mapstructure:"contentTypeHTML"`
	ContentTypePlain string `mapstructure:"contentTypePlain"`
}

// EmailConf 邮箱配置
var EmailConf = &email{}

// database 数据库配置结构
type database struct {
	DBType      string `mapstructure:"dbType"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	DBName      string `mapstructure:"dbName"`
	TablePrefix string `mapstructure:"tablePrefix"`
	Debug       bool   `mapstructure:"debug"`
}

// DBConf 数据库配置
var DBConf = &database{}

// redis 缓存配置结构
type redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DBNum    int    `mapstructure:"db"`
}

// RedisConf 缓存配置
var RedisConf = &redis{}

// logger 日志配置结构
type logger struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// LoggerConf 日志配置
var LoggerConf = &logger{}

// Setup 生成服务配置
func Setup() {
	viper.SetConfigType("YAML")
	// 读取配置文件内容
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Read 'config.yaml' fail: %v\n", err)
	}
	// 配置内容解析
	viper.ReadConfig(bytes.NewBuffer(data))
	// 解析配置赋值
	viper.UnmarshalKey("server", ServerConf)
	viper.UnmarshalKey("email", EmailConf)
	viper.UnmarshalKey("database", DBConf)
	viper.UnmarshalKey("redis", RedisConf)
	viper.UnmarshalKey("logger", LoggerConf)
}
