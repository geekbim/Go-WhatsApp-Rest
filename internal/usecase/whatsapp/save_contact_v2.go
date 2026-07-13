package whatsapp

import (
	"context"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *whatsAppInteractor) SaveContactV2(ctx context.Context, jid string, msisdn string, fullName string, firstName string) *exceptions.CustomerError {
	var multierr *multierror.Error

	err := interactor.whatsAppService.WhatsAppPutContactName(ctx, jid, msisdn, fullName, firstName)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return nil
}

