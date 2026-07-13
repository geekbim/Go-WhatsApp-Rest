package whatsapp

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/valueobject"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) SendMessageV2(ctx context.Context, whatsApp *entity.WhatsApp, jid string) (*entity.WhatsApp, *exceptions.CustomerError) {
	var (
		multierr  *multierror.Error
		remoteJID types.JID
	)

	switch whatsApp.ChatType.GetValue() {
	case valueobject.Private:
		remoteJID = interactor.whatsAppService.WhatsAppComposeJID(whatsApp.Msisdn)
	case valueobject.Group:
		remoteJID = types.NewJID(whatsApp.Msisdn, types.GroupServer)
	}

	id, err := interactor.whatsAppService.WhatsAppSendText(ctx, jid, remoteJID, whatsApp.Message)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	whatsApp.Id = id

	return whatsApp, nil
}
