package entity

import (
	"errors"
	"go_wa_rest/valueobject"
	"regexp"

	"github.com/hashicorp/go-multierror"
)

type WhatsAppDocument struct {
	ChatType valueobject.ChatType
	Msisdn   string
	Document []byte
	FileName string
	FileType string
}

type WhatsAppDocumentDTO struct {
	ChatType valueobject.ChatTypeEnum
	Msisdn   string
	Document []byte
	FileName string
	FileType string
}

func NewWhatsAppDocument(whatsAppDocumentDTO *WhatsAppDocumentDTO) (*WhatsAppDocument, *multierror.Error) {
	var multierr *multierror.Error

	whatsappDocument := &WhatsAppDocument{
		ChatType: valueobject.NewChatType(whatsAppDocumentDTO.ChatType),
		Msisdn:   whatsAppDocumentDTO.Msisdn,
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

	if w.ChatType.GetValue() == "" {
		multierr = multierror.Append(multierr, errors.New("chat type cannot be empty"))
	}

	switch w.ChatType.GetValue() {
	case valueobject.Private:
		re := regexp.MustCompile(`(0|\+62|062|62)[0-9]+$`)

		if !re.MatchString(w.Msisdn) {
			multierr = multierror.Append(multierr, errors.New("invalid msisdn"))
		}
	case valueobject.Group:
		if w.Msisdn == "" {
			multierr = multierror.Append(multierr, errors.New("invalid msisdn"))
		}
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
