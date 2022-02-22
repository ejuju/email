// Package email provides utilities to connect to an SMTP server via TLS and send emails
// The goal of this package is to extend the smtp and mail packages from the Go standard library
//
// This package helps you do 3 things:
// 1. Connecting to a SMTP server (using func email.NewClient)
// 2. Generating an email message string from a structured data (using struct email.Email and func email.ToMessageStr)
// 3. Sending emails over the same connection to the SMTP server (using func email.Send)
//
// Tested with Mailjet, not sure how well it works with other email providers
package email
