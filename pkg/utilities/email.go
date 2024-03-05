package utilities

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Host        string
	Port        int
	Sender      string
	Email       string
	AppPassword string
}

func GetSMTPConfig() SMTPConfig {
	port, _ := strconv.ParseInt(os.Getenv("CONFIG_SMTP_PORT"), 10, 64)
	return SMTPConfig{
		Host:        os.Getenv("CONFIG_SMTP_HOST"),
		Port:        int(port),
		Sender:      os.Getenv("CONFIG_SENDER_NAME"),
		Email:       os.Getenv("CONFIG_AUTH_EMAIL"),
		AppPassword: os.Getenv("CONFIG_AUTH_PASSWORD"),
	}
}

func SendOTPToEmail(otp string, target string) error {
	env := GetSMTPConfig()
	mailBody := fmt.Sprint("OTP: ", otp)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", env.Sender)
	mailer.SetHeader("To", target)
	mailer.SetHeader("Subject", "Superstore Admin Invitation")
	mailer.SetBody("text/html", mailBody)

	dialer := gomail.NewDialer(
		env.Host,
		env.Port,
		env.Email,
		env.AppPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Mail sent!")
	return nil
}
