package domain

import "github.com/aws/aws-sdk-go/service/sesv2"

type SmtpService interface {
	SendEmail(input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error)
}
