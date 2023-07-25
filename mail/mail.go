package mail

import (
	"bytes"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
)

func SendMail(subject, body string) error {
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	from := mail.Address{Name: "RSS Summarize", Address: email}
	to := mail.Address{Name: "", Address: email}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = subject

	var msg bytes.Buffer
	for k, v := range header {
		msg.WriteString(k)
		msg.WriteString(": ")
		msg.WriteString(v)
		msg.WriteString("\r\n")
	}

	msg.WriteString("\r\n" + body)

	auth := smtp.PlainAuth("", from.Address, password, host)

	err := smtp.SendMail(address, auth, from.Address, []string{to.Address}, msg.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}

	return nil
}
