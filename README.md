# Simple email utilities for Go

## Providers

Works with:

- Mailjet (SMTP relay)

Doesn't seem to work with:

- Gmail (smtp.google.com:465 and smtp.google.com:587)

---

## Examples

### Sending an email

```go
import "github.com/ejuju/email"

func main() {
	// init smtp client
	smtpClient, err := NewClient(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)
	if err != nil {
		// handle error
		return
	}

	// send email
	err = Send(smtpClient, Email{
		From: mail.Address{
			Name:    os.Getenv("EMAIL_NAME"),
			Address: os.Getenv("EMAIL_ADDR"),
		},
		To: []mail.Address{
			mail.Address{
				Address: "test1@test.com",
			},
		},
		Subject: "This is a subject",
		Body: Body{
			HTML:      "<html><h1>This is the HTML body</h1></html>",
			PlainText: "This is a plain text fallback",
		},
	})
	if err != nil {
		// handle error
		return
	}

	// disconnect client
	err = smtpClient.Quit()
	if err != nil {
		// handle error
		return
	}
}
```

### Sending multiple emails over the same connection

```go
import "github.com/ejuju/email"

func main() {
	// init smtp client
	smtpClient, err := NewClient(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)
	if err != nil {
		log.Println(err)
		return
	}

	recipients := []string{
		"test1@test.com",
		"test2@test.com",
		"test3@test.com",
	}

	// send email(s)
	for _, recipient := range recipients {
		e := Email{
			From: mail.Address{
				Name:    os.Getenv("EMAIL_NAME"),
				Address: os.Getenv("EMAIL_ADDR"),
			},
			To: []mail.Address{
				mail.Address{
					Address: recipient,
				},
			},
			Subject: "This is a subject",
			Body: Body{
				HTML:      "<html><h1>This is the HTML body</h1></html>",
				PlainText: "This is a plain text fallback",
			},
		}

		err = Send(smtpClient, e)

		if err != nil {
			log.Println(err)
			return
		}
	}

	// disconnect client
	err = smtpClient.Quit()
	if err != nil {
		log.Println(err)
		return
	}
}
```

---

## Roadmap

- Support file attachments
- Support Content-ID MIME field
