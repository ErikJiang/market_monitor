package email

import (
	"net/smtp"
	"strconv"

	"github.com/ErikJiang/market_monitor/extend/conf"
)

// SendEmail 发送邮件
func SendEmail(subject string, recvEmail string, emailContent string) error {

	auth := smtp.PlainAuth(
		"",
		conf.EmailConf.UserName,
		conf.EmailConf.Password,
		conf.EmailConf.Host,
	)

	msg := []byte(
		"To: " + recvEmail + "\r\n" +
			"From: " + conf.EmailConf.ServName + "<" + conf.EmailConf.UserName + ">\r\n" +
			"Subject: " + subject + "\r\n" + "MIME-version: 1.0;\nContent-Type: " +
			conf.EmailConf.ContentTypeHTML + ";charset=\"UTF-8\";\t\n\r\n" + emailContent,
	)

	err := smtp.SendMail(
		conf.EmailConf.Host+":"+strconv.Itoa(conf.EmailConf.Port),
		auth,
		conf.EmailConf.UserName,
		[]string{recvEmail},
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}
