package email

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to []string, msg string) {
	auth := smtp.PlainAuth(
		"",
		"islam.mohamed@gmail.com", //put your email here
		"",                        // put app pasword generated from gmail
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"islam.mohamed.devo@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
