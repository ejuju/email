package email

import (
	"crypto/tls"
	"net/smtp"
)

// NewClient returns an authenticated SMTP client or an error if the process fails
// When successfull, make sure to call Quit on the smtp client struct after you're done sending emails to close the connection to the SMTP server
func NewClient(host, username, password string) (*smtp.Client, error) {
	auth := smtp.PlainAuth("", username, password, host)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// call tls.Dial instead of smtp.Dial for smtp servers running on 465 that require an ssl connection from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", host, tlsConfig)
	if err != nil {
		return nil, err
	}

	// init smtp client
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}

	// authenticate
	err = c.Auth(auth)
	if err != nil {
		return nil, err
	}

	return c, nil
}
