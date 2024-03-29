package whatsapp

import (
	"context"
	"go_wa_rest/internal/service/whatsapp"
	"go_wa_rest/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (interactor *whatsAppInteractor) GetQr(ctx context.Context) (string, int, *exceptions.CustomerError) {
	var (
		qrImage   string
		qrTimeOut int
		multierr  *multierror.Error
	)

	interactor.waClient = whatsapp.InitWhatsApp()

	if interactor.waClient.Store.ID == nil {
		// No ID stored, new login
		qrChan, err := interactor.waClient.GetQRChannel(ctx)
		if err != nil {
			multierr = multierror.Append(multierr, err)
			return "", 0, &exceptions.CustomerError{
				Status: exceptions.ERRBUSSINESS,
				Errors: multierr,
			}
		}
		err = interactor.waClient.Connect()
		if err != nil {
			panic(err)
		}

		interactor.waClient.SendPresence("available")

		qrImage, qrTimeOut = interactor.whatsAppService.WhatsAppGenerateQR(qrChan)
		qrImage = "data:image/png;base64," + qrImage
	} else {
		// Already logged in, just connect
		err := interactor.waClient.Connect()
		if err != nil {
			multierr = multierror.Append(multierr, err)
			return "", 0, &exceptions.CustomerError{
				Status: exceptions.ERRBUSSINESS,
				Errors: multierr,
			}
		}
		interactor.waClient.SendPresence("available")
	}

	return qrImage, qrTimeOut, nil
}
