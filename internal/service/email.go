package service

import (
	"errors"
	"bytes"
	"html/template"
	"path/filepath"
	"crypto/tls"
	"log"

	"github.com/go-gomail/gomail"
)

var tempalteByType map[string]string = map[string]string{
	"DebtExceeded": "email_debt_exceeded.html",
	"TaskResult": "email_task_result.html",
}

type EmailService struct {
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
}

func NewEmailService(host, username, password string, port int) *EmailService {
	return &EmailService{
		smtpHost: host,
		smtpPort: port,
		smtpUsername: username,
		smtpPassword: password,
	}
}

func (s *EmailService) sendEmail(toEmail, subject, body string) error {
	email := gomail.NewMessage()
	email.SetHeader("From", s.smtpUsername)
	email.SetHeader("To", toEmail)
	email.SetHeader("Subject", subject)
	email.SetBody("text/plain", body)

	dialer := gomail.NewDialer(s.smtpHost, s.smtpPort, s.smtpUsername, s.smtpPassword)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	log.Println(body)
	err := dialer.DialAndSend(email)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	return nil
}

func (s *EmailService) SendEmail(emailType, to, subject string, data any) error {
	templateFilename, ok := tempalteByType[emailType] 
	if !ok {
		return errors.New("invalid email type provided")
	}

	templatePath := filepath.Join("templates", templateFilename)

	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = template.Execute(&body, data)
	if err != nil {
		return err
	}

	return s.sendEmail(to, subject, body.String())
}