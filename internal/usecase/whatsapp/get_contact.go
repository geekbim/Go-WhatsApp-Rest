package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) GetContact(ctx context.Context) (map[types.JID]types.ContactInfo, *exceptions.CustomerError) {
	var multierr *multierror.Error

	if interactor.waClient == nil {
		multierr = multierror.Append(multierr, errors.New("session not found"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	contacts, err := interactor.waClient.Store.Contacts.GetAllContacts(ctx)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return contacts, nil
}
