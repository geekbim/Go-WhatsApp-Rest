package whatsapp

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/valueobject"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) SendDocumentV2(ctx context.Context, whatsAppDocument *entity.WhatsAppDocument, jid string) (*entity.WhatsAppDocument, *exceptions.CustomerError) {
	var (
		multierr  *multierror.Error
		remoteJID types.JID
	)

	switch whatsAppDocument.ChatType.GetValue() {
	case valueobject.Private:
		remoteJID = interactor.whatsAppService.WhatsAppComposeJID(whatsAppDocument.Msisdn)
	case valueobject.Group:
		remoteJID = types.NewJID(whatsAppDocument.Msisdn, types.GroupServer)
	}

	id, err := interactor.whatsAppService.WhatsAppSendDocument(ctx, jid, remoteJID, whatsAppDocument)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	whatsAppDocument.Id = id

	return whatsAppDocument, nil
}
