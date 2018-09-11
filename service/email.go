package service

import (
	"log"
	"strconv"
	"net/smtp"
	"github.com/JiangInk/market_monitor/utils"
)

type ContentTypeConf struct {
	HTML string
	Plain string
}

type EmailConf struct {
	ServiceName string
	UserName string
	Password string
	Host string
	Port int
	ContentType ContentTypeConf
}

func SendEmail(subject string, recvEmail string, emailContent string) error {
	log.Println("enter sendEmail.")

	JsonParse := utils.NewJsonStruct()
	emailConf := EmailConf{}
	JsonParse.Load("E:\\GOPRO\\src\\github.com\\JiangInk\\market_monitor\\config\\email.json", &emailConf)

	auth := smtp.PlainAuth(
		"",
		emailConf.UserName,
		emailConf.Password,
		emailConf.Host,
	)

	msg := []byte(
		"To: " + recvEmail + "\r\n" +
		"From: " + emailConf.ServiceName + "<" + emailConf.UserName + ">\r\n" +
		"Subject: " + subject + "\r\n" +
		emailConf.ContentType.HTML + "\r\n" + emailContent,
	)

	err := smtp.SendMail(
		emailConf.Host + ":" + strconv.Itoa(emailConf.Port),
		auth,
		emailConf.UserName,
		[]string{recvEmail},
		msg,
	)
	log.Println(err)
	if err != nil {
		return err
	}
	return nil
}
