package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func (interactor *whatsAppInteractor) SendMessage(ctx context.Context, whatsApp *entity.WhatsApp) (*entity.WhatsApp, *exceptions.CustomerError) {
	var multierr *multierror.Error

	remoteJID := interactor.whatsAppService.WhatsAppComposeJID(whatsApp.Msisdn)

	msgContent := &waproto.Message{
		Conversation: proto.String(whatsApp.Message),
	}

	if interactor.waClient == nil {
		multierr = multierror.Append(multierr, errors.New("session not found"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	_, err := interactor.waClient.SendMessage(ctx, remoteJID, msgContent)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return whatsApp, nil
}
