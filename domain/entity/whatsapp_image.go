package entity

import (
	"errors"
	"go_wa_rest/valueobject"
	"regexp"

	"github.com/hashicorp/go-multierror"
)

type WhatsAppImage struct {
	Id       string
	ChatType valueobject.ChatType
	Msisdn   string
	Message  string
	Image    []byte
	FileName string
	FileType string
}

type WhatsAppImageDTO struct {
	ChatType valueobject.ChatTypeEnum
	Msisdn   string
	Message  string
	Image    []byte
	FileName string
	FileType string
}

func NewWhatsAppImage(whatsAppImageDTO *WhatsAppImageDTO) (*WhatsAppImage, *multierror.Error) {
	var multierr *multierror.Error

	whatsappImage := &WhatsAppImage{
		ChatType: valueobject.NewChatType(whatsAppImageDTO.ChatType),
		Msisdn:   whatsAppImageDTO.Msisdn,
		Message:  whatsAppImageDTO.Message,
		Image:    whatsAppImageDTO.Image,
		FileName: whatsAppImageDTO.FileName,
		FileType: whatsAppImageDTO.FileType,
	}

	if errValidate := whatsappImage.Validate(); errValidate != nil {
		multierr = multierror.Append(multierr, errValidate)
	}

	if multierr != nil {
		return nil, multierr
	}

	return whatsappImage, nil
}

func (w *WhatsAppImage) Validate() *multierror.Error {
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

	if w.Message == "" {
		multierr = multierror.Append(multierr, errors.New("message cannot be empty"))
	}

	if w.Image == nil {
		multierr = multierror.Append(multierr, errors.New("image cannot be empty"))
	}

	if w.FileName == "" {
		multierr = multierror.Append(multierr, errors.New("file name cannot be empty"))
	}

	if w.FileType == "" {
		multierr = multierror.Append(multierr, errors.New("file type cannot be empty"))
	}

	return multierr
}
