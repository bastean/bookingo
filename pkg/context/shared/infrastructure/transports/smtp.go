package transports

import (
	"net/smtp"
)

type SMTP struct {
	smtp.Auth
	MIMEHeaders                        string
	SMTPServerURL, Hotelname, Password string
	ServerURL                          string
}

func NewSMTP(host, port, hotelname, password, serverURL string) *SMTP {
	return &SMTP{
		Auth:          smtp.PlainAuth("", hotelname, password, host),
		MIMEHeaders:   "MIME-version: 1.0;\n" + "Content-Type: text/html; charset=\"UTF-8\";\n\n",
		SMTPServerURL: host + ":" + port,
		Hotelname:     hotelname,
		Password:      password,
		ServerURL:     serverURL,
	}
}
