package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *whatsAppInteractor) Logout(ctx context.Context) *exceptions.CustomerError {
	var multierr *multierror.Error

	if interactor.waClient.Store.ID != nil {
		interactor.waClient.SendPresence("unavailable")
		if err := interactor.waClient.Logout(); err != nil {
			interactor.waClient.Disconnect()
			if err := interactor.waClient.Store.Delete(); err != nil {
				multierr = multierror.Append(multierr, err)
				return &exceptions.CustomerError{
					Status: exceptions.ERRBUSSINESS,
					Errors: multierr,
				}
			}
		}

		interactor.waClient = nil

		return nil
	}

	multierr = multierror.Append(multierr, errors.New("whatsApp client is not valid"))
	return &exceptions.CustomerError{
		Status: exceptions.ERRBUSSINESS,
		Errors: multierr,
	}
}
