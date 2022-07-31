package entity

import (
	"errors"
	"regexp"

	"github.com/hashicorp/go-multierror"
)

type WhatsApp struct {
	Msisdn  string
	Message string
}

type WhatsAppDTO struct {
	Msisdn  string
	Message string
}

func NewWhatsApp(whatsAppDTO *WhatsAppDTO) (*WhatsApp, *multierror.Error) {
	var multierr *multierror.Error

	whatsapp := &WhatsApp{
		Msisdn:  whatsAppDTO.Msisdn,
		Message: whatsAppDTO.Message,
	}

	if errValidate := whatsapp.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return whatsapp, nil
}

func (w *WhatsApp) Validate() *multierror.Error {
	var multierr *multierror.Error

	re := regexp.MustCompile(`(0|\+62|062|62)[0-9]+$`)

	if !re.MatchString(w.Msisdn) {
		multierr = multierror.Append(multierr, errors.New("invalid msisdn"))
	}

	if w.Message == "" {
		multierr = multierror.Append(multierr, errors.New("message cannot be empty"))
	}

	return multierr
}
