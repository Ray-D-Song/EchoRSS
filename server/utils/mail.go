package utils

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body string) {
	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

}
