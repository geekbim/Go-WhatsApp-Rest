package testdata

import (
	"go_wa_rest/domain/entity"
	"go_wa_rest/valueobject"
)

func NewWhatsAppDocumentDTO() *entity.WhatsAppDocumentDTO {
	return &entity.WhatsAppDocumentDTO{
		ChatType: valueobject.Group,
		Msisdn:   "622150942316",
		Message:  "hello",
		Document: []byte{},
		FileName: "test",
		FileType: "document/pdf",
	}
}

func NewWhatsAppDocument(whatsAppDocumentDTO *entity.WhatsAppDocumentDTO) *entity.WhatsAppDocument {
	return &entity.WhatsAppDocument{
		ChatType: valueobject.NewChatType(whatsAppDocumentDTO.ChatType),
		Msisdn:   whatsAppDocumentDTO.Msisdn,
		Message:  whatsAppDocumentDTO.Message,
		Document: whatsAppDocumentDTO.Document,
		FileName: whatsAppDocumentDTO.FileName,
		FileType: whatsAppDocumentDTO.FileType,
	}
}
