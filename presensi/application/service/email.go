package service

type EmailService interface {
	SendEmail(to string, from string, subject string)
}
