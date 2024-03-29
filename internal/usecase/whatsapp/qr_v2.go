package whatsapp

import (
	"context"
	whatsapp_service "go_wa_rest/internal/service/whatsapp"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *whatsAppInteractor) GetQrV2(ctx context.Context, jid string) (string, int, *exceptions.CustomerError) {
	var multierr *multierror.Error

	whatsapp_service.InitWhatsAppV2(nil, jid)

	qrCodeImage, qrCodeTimeout, err := interactor.whatsAppService.WhatsAppLogin(jid)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return "", 0, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	// If Return is Not QR Code But Reconnected
	// Then Return OK With Reconnected Status
	if qrCodeImage == "WhatsApp Client is Reconnected" {
		return qrCodeImage, qrCodeTimeout, nil
	}

	return qrCodeImage, qrCodeTimeout, nil
}
