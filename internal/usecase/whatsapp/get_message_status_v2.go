package whatsapp

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *whatsAppInteractor) GetMessageStatusV2(ctx context.Context, whatsAppStatus *entity.WhatsAppStatus, jid string) (*entity.WhatsAppStatus, *exceptions.CustomerError) {
	var multierr *multierror.Error

	res, err := interactor.whatsAppService.WhatsAppMessageStatus(ctx, jid, whatsAppStatus.MessageId)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return res, nil
}
