package whatsapp

import (
	"context"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) GetContactV2(ctx context.Context, jid string) (map[types.JID]types.ContactInfo, *exceptions.CustomerError) {
	var multierr *multierror.Error

	contacts, err := interactor.whatsAppService.WhatsAppContact(jid)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return contacts, nil
}
