package testdata

import (
	"go_wa_rest/domain/entity"
	"go_wa_rest/valueobject"
)

func NewWhatsAppImageDTO() *entity.WhatsAppImageDTO {
	return &entity.WhatsAppImageDTO{
		ChatType: valueobject.Group,
		Msisdn:   "622150942316",
		Image:    []byte{},
		FileName: "test",
		FileType: "image/jpg",
	}
}

func NewWhatsAppImage(whatsAppImageDTO *entity.WhatsAppImageDTO) *entity.WhatsAppImage {
	return &entity.WhatsAppImage{
		ChatType: valueobject.NewChatType(whatsAppImageDTO.ChatType),
		Msisdn:   whatsAppImageDTO.Msisdn,
		Image:    whatsAppImageDTO.Image,
		FileName: whatsAppImageDTO.FileName,
		FileType: whatsAppImageDTO.FileType,
	}
}
