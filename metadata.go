package email

// Utilities for email metadata (for header fields and body part data)

// Content type values
const (
	TextPlain            = "text/plain; " + CharsetUTF8
	TextHTML             = "text/html; " + CharsetUTF8
	CharsetUTF8          = "charset=\"UTF-8\""
	MultipartAlternative = "multipart/alternative; boundary=\"" + BoundaryVal + "\""
	BoundaryVal          = "part"             // the boundary value chosen here is arbitrary but should be consistent with the email's alternative body parts
	QuotedPrintable      = "quoted-printable" //
)
