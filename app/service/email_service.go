package service

import (
	"github.com/phuslu/log"
	"net/smtp"
)

func sendEmail(from string, password string, to []string, subject string, body string) error {
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, []byte(subject+body))
	if err != nil {
		log.Error().Err(err).Msg("smtp.SendMail failed")
		return err
	}
	return nil
}
