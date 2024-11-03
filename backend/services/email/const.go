package email

type EmailInput struct {
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Message string
}
