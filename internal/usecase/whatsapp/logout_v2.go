package whatsapp

import (
	"context"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *whatsAppInteractor) LogoutV2(ctx context.Context, jid string) *exceptions.CustomerError {
	var multierr *multierror.Error

	err := interactor.whatsAppService.WhatsAppLogout(jid)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return nil
}
