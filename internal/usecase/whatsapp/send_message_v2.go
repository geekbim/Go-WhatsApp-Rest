package whatsapp

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/pkg/service/whatsapp"

	"github.com/hashicorp/go-multierror"
)

func (interactor *whatsAppInteractor) SendMessageV2(ctx context.Context, whatsApp *entity.WhatsApp, jid string) (*entity.WhatsApp, *exceptions.CustomerError) {
	var multierr *multierror.Error

	_, err := whatsapp.WhatsAppSendText(ctx, jid, whatsApp.Msisdn, whatsApp.Message)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return whatsApp, nil
}
