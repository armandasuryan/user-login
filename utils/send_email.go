package utils

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type SentMail struct {
	log *logrus.Logger
}

func (s *SentMail) sentEmail(headerFrom string, headerTo string, subject string, body string) error {
	s.log.Println("Execute function sent email")

	m := gomail.NewMessage()
	m.SetHeader("From", headerFrom)
	m.SetHeader("To", headerTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// smtp config
	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	user := os.Getenv("SMTP_USER")
	passwd := os.Getenv("SMTP_PASSWORD")

	configEmail := gomail.NewDialer(host, port, user, passwd)

	if err := configEmail.DialAndSend(m); err != nil {
		s.log.Error("Failed to sent email")
		return err
	}

	s.log.Println("Email sent successfully!")

	return nil
}
