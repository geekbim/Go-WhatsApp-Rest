package whatsapp

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/valueobject"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) SendImageV2(ctx context.Context, whatsAppImage *entity.WhatsAppImage, jid string) (*entity.WhatsAppImage, *exceptions.CustomerError) {
	var (
		multierr  *multierror.Error
		remoteJID types.JID
	)

	switch whatsAppImage.ChatType.GetValue() {
	case valueobject.Private:
		remoteJID = interactor.whatsAppService.WhatsAppComposeJID(whatsAppImage.Msisdn)
	case valueobject.Group:
		remoteJID = types.NewJID(whatsAppImage.Msisdn, types.GroupServer)
	}

	_, err := interactor.whatsAppService.WhatsAppSendImage(ctx, jid, remoteJID, whatsAppImage)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return whatsAppImage, nil
}
