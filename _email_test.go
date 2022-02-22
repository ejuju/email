// package email

// import (
// 	"net/mail"
// 	"os"
// 	"testing"

// 	"github.com/ejuju/dotenv"
// )

// func TestMessage(t *testing.T) {
// 	// inject test variables from .env file
// 	err := dotenv.Load("./.env")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	// init smtp client
// 	smtpClient, err := NewClient(
// 		os.Getenv("SMTP_HOST"),
// 		os.Getenv("SMTP_USERNAME"),
// 		os.Getenv("SMTP_PASSWORD"),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	err = Send(
// 		smtpClient,
// 		Email{
// 			From: mail.Address{
// 				Name:    os.Getenv("EMAIL_NAME"),
// 				Address: os.Getenv("EMAIL_ADDR"),
// 			},
// 			To: []mail.Address{
// 				mail.Address{
// 					Address: os.Getenv("EMAIL_RECIPIENT_2"),
// 				},
// 			},
// 			CC: []mail.Address{
// 				mail.Address{
// 					Address: os.Getenv("EMAIL_RECIPIENT_3"),
// 				},
// 			},
// 			BCC: []mail.Address{
// 				mail.Address{
// 					Address: os.Getenv("EMAIL_RECIPIENT"),
// 				},
// 			},
// 			ReplyTo: mail.Address{
// 				Address: os.Getenv("EMAIL_REPLY_TO"),
// 			},
// 			Subject: "7nd test email subject",
// 			Body: Body{
// 				PlainText: "Test plain text body",
// 				HTML:      "<html><h1>Test html body</h1></html>",
// 			},
// 		},
// 	)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	err = smtpClient.Quit()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// }
