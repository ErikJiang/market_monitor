package email

import (
	"net/smtp"
	"strconv"

	"github.com/JiangInk/market_monitor/config"
)

func SendEmail(subject string, recvEmail string, emailContent string) error {

	auth := smtp.PlainAuth(
		"",
		config.EmailSetting.UserName,
		config.EmailSetting.Password,
		config.EmailSetting.Host,
	)

	msg := []byte(
		"To: " + recvEmail + "\r\n" +
			"From: " + config.EmailSetting.ServName + "<" + config.EmailSetting.UserName + ">\r\n" +
			"Subject: " + subject + "\r\n" + "MIME-version: 1.0;\nContent-Type: " +
			config.EmailSetting.ContentTypeHTML + ";charset=\"UTF-8\";\t\n\r\n" + emailContent,
	)

	err := smtp.SendMail(
		config.EmailSetting.Host+":"+strconv.Itoa(config.EmailSetting.Port),
		auth,
		config.EmailSetting.UserName,
		[]string{recvEmail},
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}
