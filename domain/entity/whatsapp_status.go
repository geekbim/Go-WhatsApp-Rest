package entity

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

type WhatsAppStatus struct {
	MessageId string
	Status    string
}

type WhatsAppStatusDTO struct {
	MessageId string
}

func NewWhatsAppStatus(whatsAppStatusDTO *WhatsAppStatusDTO) (*WhatsAppStatus, *multierror.Error) {
	var multierr *multierror.Error

	whatsappStatus := &WhatsAppStatus{
		MessageId: whatsAppStatusDTO.MessageId,
	}

	if errValidate := whatsappStatus.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return whatsappStatus, nil
}

func (w *WhatsAppStatus) Validate() *multierror.Error {
	var multierr *multierror.Error

	if w.MessageId == "" {
		multierr = multierror.Append(multierr, errors.New("message id cannot be empty"))
	}

	return multierr
}
