package email

import (
	"net/smtp"
	"strconv"

	"github.com/JiangInk/market_monitor/config"
)

// SendEmail 发送邮件
func SendEmail(subject string, recvEmail string, emailContent string) error {

	auth := smtp.PlainAuth(
		"",
		config.EmailConf.UserName,
		config.EmailConf.Password,
		config.EmailConf.Host,
	)

	msg := []byte(
		"To: " + recvEmail + "\r\n" +
			"From: " + config.EmailConf.ServName + "<" + config.EmailConf.UserName + ">\r\n" +
			"Subject: " + subject + "\r\n" + "MIME-version: 1.0;\nContent-Type: " +
			config.EmailConf.ContentTypeHTML + ";charset=\"UTF-8\";\t\n\r\n" + emailContent,
	)

	err := smtp.SendMail(
		config.EmailConf.Host+":"+strconv.Itoa(config.EmailConf.Port),
		auth,
		config.EmailConf.UserName,
		[]string{recvEmail},
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}
