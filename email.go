package email

import (
	"net/mail"
	"strings"
)

// Email holds necessary information to build an email message (= header fields + body)
type Email struct {
	From       mail.Address
	ReplyTo    mail.Address
	ReturnPath mail.Address
	To         []mail.Address
	CC         []mail.Address
	BCC        []mail.Address
	Subject    string
	Body       Body
}

// Body holds the parts that make up the email body
type Body struct {
	PlainText string
	HTML      string
}

// ToMessageStr generates the message string that will be sent to the SMTP server
func ToMessageStr(e Email) string {
	headerMap := map[string]string{
		"From":                      e.From.String(),
		"Reply-To":                  e.ReplyTo.String(),
		"To":                        mailAddressesToStr(e.To),
		"Cc":                        mailAddressesToStr(e.CC),
		"Bcc":                       mailAddressesToStr(e.BCC),
		"Subject":                   e.Subject,
		"MIME-Version":              "1.0",
		"Content-Type":              getEmailContentType(e.Body),
		"Content-Transfer-Encoding": getEmailContentTransferEncoding(e.Body),
	}
	header := ""
	for key, val := range headerMap {
		header += key + ":" + val + "\r\n"
	}

	body := emailBodyToStr(e.Body, BoundaryVal)

	return header + "\r\n" + body + "\r\n"
}

func getEmailContentType(body Body) string {
	if body.HTML != "" {
		if body.PlainText != "" {
			// plain text and html are defined => multipart/alternative
			return MultipartAlternative
		}
		// only html => text/html
		return TextHTML
	}

	// no html, only plain text / or even if no plain text => text/plain
	return TextPlain
}

func getEmailContentTransferEncoding(body Body) string {
	return QuotedPrintable
}

func mailAddressesToStr(addrs []mail.Address) string {
	var list []string
	for _, a := range addrs {
		if a.Address == "" {
			continue
		}
		list = append(list, a.String())
	}
	return strings.Join(list, "; ")
}

func emailBodyToStr(body Body, boundary string) string {
	var out string
	out += bodyPartToStr(body.PlainText, boundary, TextPlain, QuotedPrintable)
	out += bodyPartToStr(body.HTML, boundary, TextHTML, QuotedPrintable)
	out += closeBodyParts(boundary)
	return out
}

func bodyPartToStr(content string, boundary string, contentType string, contentTransferEncoding string) string {
	if content == "" {
		return ""
	}
	out := "--" + boundary + "\n"                                           // boundary
	out += "Content-Type: " + contentType + "\n"                            // content type
	out += "Content-Transfer-Encoding: " + contentTransferEncoding + "\n\n" // content transfer encoding
	out += content + "\n\n"                                                 // content
	return out
}

func closeBodyParts(boundary string) string {
	return "--" + boundary + "--" // line to end multipart body
}
