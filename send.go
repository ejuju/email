package email

import "net/smtp"

// Send sends the email to the client without calling client.Quit() so that the connection can be used to send several emails
// A call to Send should be preceded by a call to NewClient() to connect to the server and authenticate
// A call to Send should ultimately be followed by a call to client.Quit() so that the SMTP client can close the server connection
func Send(client *smtp.Client, e Email) error {
	// initiate new mail
	err := client.Mail(e.From.Address)
	if err != nil {
		return err
	}

	// set "To" recipients
	for _, r := range e.To {
		err = client.Rcpt(r.Address)
		if err != nil {
			return err
		}
	}

	// set "Cc" recipients
	for _, r := range e.CC {
		err = client.Rcpt(r.Address)
		if err != nil {
			return err
		}
	}

	// set "Bcc" recipients
	for _, r := range e.BCC {
		err = client.Rcpt(r.Address)
		if err != nil {
			return err
		}
	}

	// get data writer
	w, err := client.Data()
	if err != nil {
		return err
	}

	// write content
	_, err = w.Write([]byte(ToMessageStr(e)))
	if err != nil {
		return err
	}

	// close writer
	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}
