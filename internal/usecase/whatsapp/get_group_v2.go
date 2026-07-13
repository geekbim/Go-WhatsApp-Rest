package whatsapp

import (
	"context"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) GetGroupV2(ctx context.Context, jid string) ([]*types.GroupInfo, *exceptions.CustomerError) {
	var multierr *multierror.Error

	groups, err := interactor.whatsAppService.WhatsAppGroup(jid)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return groups, nil
}
