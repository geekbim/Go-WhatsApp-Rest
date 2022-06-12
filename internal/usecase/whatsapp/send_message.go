package whatsapp

import (
	"context"
	"go-wa-rest/domain/entity"
	"go-wa-rest/pkg/exceptions"
	"go-wa-rest/pkg/service/whatsapp"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func (interactor *whatsAppInteractor) SendMessage(ctx context.Context, whatsApp *entity.WhatsApp) (*entity.WhatsApp, *exceptions.CustomerError) {
	var multierr *multierror.Error

	remoteJID := whatsapp.WhatsAppComposeJID(whatsApp.Msisdn)

	msgId := whatsmeow.GenerateMessageID()
	msgContent := &waproto.Message{
		Conversation: proto.String(whatsApp.Message),
	}

	_, err := interactor.waClient.SendMessage(remoteJID, msgId, msgContent)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return whatsApp, nil
}
