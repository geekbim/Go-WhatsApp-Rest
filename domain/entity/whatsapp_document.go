package entity

import (
	"errors"
	"regexp"

	"github.com/hashicorp/go-multierror"
)

type WhatsAppDocument struct {
	Msisdn   string
	Message  string
	Document []byte
	FileName string
	FileType string
}

type WhatsAppDocumentDTO struct {
	Msisdn   string
	Message  string
	Document []byte
	FileName string
	FileType string
}

func NewWhatsAppDocument(whatsAppDocumentDTO *WhatsAppDocumentDTO) (*WhatsAppDocument, *multierror.Error) {
	var multierr *multierror.Error

	whatsappDocument := &WhatsAppDocument{
		Msisdn:   whatsAppDocumentDTO.Msisdn,
		Message:  whatsAppDocumentDTO.Message,
		Document: whatsAppDocumentDTO.Document,
		FileName: whatsAppDocumentDTO.FileName,
		FileType: whatsAppDocumentDTO.FileType,
	}

	if errValidate := whatsappDocument.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return whatsappDocument, nil
}

func (w *WhatsAppDocument) Validate() *multierror.Error {
	var multierr *multierror.Error

	re := regexp.MustCompile(`(0|\+62|062|62)[0-9]+$`)

	if !re.MatchString(w.Msisdn) {
		multierr = multierror.Append(multierr, errors.New("invalid msisdn"))
	}

	if w.Message == "" {
		multierr = multierror.Append(multierr, errors.New("message cannot be empty"))
	}

	if w.Document == nil {
		multierr = multierror.Append(multierr, errors.New("document cannot be empty"))
	}

	if w.FileName == "" {
		multierr = multierror.Append(multierr, errors.New("file name cannot be empty"))
	}

	if w.FileType == "" {
		multierr = multierror.Append(multierr, errors.New("file type cannot be empty"))
	}

	return multierr
}
