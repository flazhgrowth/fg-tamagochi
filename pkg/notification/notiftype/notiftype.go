package notiftype

import "github.com/getbrevo/brevo-go/lib"

type (
	NotificationType string
	Driver           string

	Config struct {
		Driver Driver
		APIKey string
	}
)

var (
	NotificationEmail NotificationType = "email"

	DriverBrevo Driver = "brevo"
)

type SendArgs struct {
	Type        NotificationType
	TemplateID  int64
	Subject     string
	SenderEmail string
	SenderName  string
	To          []string
	Cc          []string
	Bcc         []string
	DataBinding map[string]any
}

type SendResponse struct {
	MessageID  string
	Successful bool
}

func (args *SendArgs) GetTo() []lib.SendSmtpEmailTo {
	res := []lib.SendSmtpEmailTo{}
	for _, to := range args.To {
		res = append(res, lib.SendSmtpEmailTo{Email: to})
	}

	return res
}

func (args *SendArgs) GetCc() []lib.SendSmtpEmailCc {
	res := []lib.SendSmtpEmailCc{}
	for _, to := range args.Cc {
		res = append(res, lib.SendSmtpEmailCc{Email: to})
	}

	return res
}

func (args *SendArgs) GetBcc() []lib.SendSmtpEmailBcc {
	res := []lib.SendSmtpEmailBcc{}
	for _, to := range args.Bcc {
		res = append(res, lib.SendSmtpEmailBcc{Email: to})
	}

	return res
}
