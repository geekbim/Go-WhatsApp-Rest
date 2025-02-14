package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) GetGroup(ctx context.Context) ([]*types.GroupInfo, *exceptions.CustomerError) {
	var multierr *multierror.Error

	if interactor.waClient == nil {
		multierr = multierror.Append(multierr, errors.New("session not found"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	groups, err := interactor.waClient.GetJoinedGroups()
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return groups, nil
}
