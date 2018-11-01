package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"github.com/spf13/viper"
)

// Server1 测试
type Server1 struct {
	RunMode  string `mapstructure:"runMode"`
	Port int `mapstructure:"port"`
}

// Server1Setting 测试
var Server1Setting = &Server1{}


// Setup1 生成服务配置
func Setup1() {
	viper.SetConfigType("YAML")
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Read 'config.yaml' fail: %v\n", err)
	}
	// log.Println("data: ", string(data))
	viper.ReadConfig(bytes.NewBuffer(data))
	
	// server := viper.GetStringMap("server")
	// email:= viper.GetStringMap("email")
	// database:= viper.GetStringMap("database")
	// redis:= viper.GetStringMap("redis")
	// logger:= viper.GetStringMap("logger")

	viper.UnmarshalKey("server", Server1Setting)
	// 配置变动重载 todo
	log.Printf("server: %v\n", Server1Setting)
	log.Printf("server.RunMode: %v\n", Server1Setting.RunMode)
	log.Printf("server.Port: %v\n", Server1Setting.Port)

}
