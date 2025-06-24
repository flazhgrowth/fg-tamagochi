package brevo

import (
	"context"

	"github.com/flazhgrowth/fg-tamagochi/pkg/notification/notiftype"
	"github.com/getbrevo/brevo-go/lib"
)

func (impl *BrevoImpl) Send(ctx context.Context, args notiftype.SendArgs) (resp *notiftype.SendResponse, err error) {
	sendTransacEmailArgs := lib.SendSmtpEmail{
		TemplateId: args.TemplateID,
		Sender: &lib.SendSmtpEmailSender{
			Email: args.SenderEmail,
			Name:  args.SenderName,
		},
		Subject: args.Subject,
		To:      args.GetTo(),
		Params:  args.DataBinding,
	}
	cc, bcc := args.GetCc(), args.GetBcc()
	if len(cc) > 0 {
		sendTransacEmailArgs.Cc = cc
	}
	if len(bcc) > 0 {
		sendTransacEmailArgs.Bcc = bcc
	}
	li, sendResp, err := impl.client.TransactionalEmailsApi.SendTransacEmail(ctx, sendTransacEmailArgs)
	if err != nil {
		return nil, err
	}

	return &notiftype.SendResponse{
		Successful: sendResp.StatusCode < 300,
		MessageID:  li.MessageId,
	}, nil
}
