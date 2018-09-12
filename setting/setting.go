package setting

import (
	"log"
	"github.com/go-ini/ini"
)

type Server struct {
	RunMode string
	HttpPort int
}

var ServerSetting = &Server{}

type Email struct {
	ServName string
	UserName string
	Password string
	Host string
	Port int
	ContentTypeHTML string
	ContentTypePlain string
}

var EmailSetting = &Email{}

type Database struct {
	DBType string
	User string
	Password string
	Host string
	Port int
	DBName string
	TablePrefix string
}

var DBSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v\n", err)
	}

	mapTo("server", ServerSetting)
	mapTo("email", EmailSetting)
	mapTo("database", DBSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.MapTo %s setting err: %v\n", section, err)
	}
}
