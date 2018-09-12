package service

import (
	"log"
	"strconv"
	"net/smtp"
	"github.com/JiangInk/market_monitor/setting"
)

func SendEmail(subject string, recvEmail string, emailContent string) error {
	log.Println("enter sendEmail.")

	auth := smtp.PlainAuth(
		"",
		setting.EmailSetting.UserName,
		setting.EmailSetting.Password,
		setting.EmailSetting.Host,
	)

	msg := []byte(
		"To: " + recvEmail + "\r\n" +
		"From: " + setting.EmailSetting.ServName + "<" + setting.EmailSetting.UserName + ">\r\n" +
		"Subject: " + subject + "\r\n" + "MIME-version: 1.0;\nContent-Type: " + 
		setting.EmailSetting.ContentTypeHTML + ";charset=\"UTF-8\";\t\n\r\n" + emailContent,
	)

	err := smtp.SendMail(
		setting.EmailSetting.Host + ":" + strconv.Itoa(setting.EmailSetting.Port),
		auth,
		setting.EmailSetting.UserName,
		[]string{recvEmail},
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}
