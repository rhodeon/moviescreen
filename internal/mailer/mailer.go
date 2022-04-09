package mailer

import (
	"bytes"
	"embed"
	"github.com/go-mail/mail/v2"
	"html/template"
	"time"
)

//go:embed "templates"
var templateFS embed.FS

// The Mailer struct contains a dialer instance and the sender information for sending emails.
type Mailer struct {
	dialer *mail.Dialer
	sender string
}

// New returns a Mailer pointer with the given SMTP settings.
func New(host string, port int, username string, password string, sender string) *Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 5 * time.Second

	return &Mailer{
		dialer: dialer,
		sender: sender,
	}
}

// Send takes the recipient email address, the name of the email template file, and any dynamic data for the template.
func (m *Mailer) Send(recipient string, templateFile string, data any) error {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	// attempt sending mail 3 times
	for i := 1; i <= 3; i++ {
		err = m.dialer.DialAndSend(msg)
		// return if no error occurs
		if err == nil {
			return nil
		}

		// make next attempt on failure with a delay of 500 milliseconds
		time.Sleep(500 * time.Millisecond)
	}

	return err
}
