package otp

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"shool/pkg/email"
)

func SendOTP(em string, otp string) {

	var data bytes.Buffer

	temp, err := template.ParseFiles("pkg/otp/otp.html")
	if err != nil {
		fmt.Println("Errorororororororrorororororo")
		log.Fatal(err.Error())
	}
	err = temp.Execute(&data, struct {
		UserMail string
		Otp      string
	}{UserMail: em, Otp: otp})
	if err != nil {
		log.Fatal(err.Error())
	}

	body := data.String()

	header := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: Wellcom To School" + "\n" + header + "\n\n" + body

	email.SendEmail([]string{em}, msg)
}
