package query

import (
	"fmt"
	_ "github.com/codyguo/godaemon"
	"gopkg.in/gomail.v2"
)

func SendEmail(conf Config, text string) {
	if !conf.NeedEmail {
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", conf.FromEmail)
	m.SetHeader("To", conf.ToEmail)
	m.SetHeader("Subject", fmt.Sprintf("%v %v\n", conf.EmailSubject, text))

	m.SetBody("text/html", fmt.Sprintf("%v %v\n", conf.EmailSubject, text))

	d := gomail.NewDialer(conf.Smtp, conf.Port, conf.FromEmail, conf.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("err Dail And Send:%v\n", err)
	}
}
