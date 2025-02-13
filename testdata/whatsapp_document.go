package testdata

import "go_wa_rest/domain/entity"

func NewWhatsAppDocumentDTO() *entity.WhatsAppDocumentDTO {
	return &entity.WhatsAppDocumentDTO{
		Msisdn:   "622150942316",
		Message:  "hello",
		Document: []byte{},
		FileName: "test",
		FileType: "document/pdf",
	}
}

func NewWhatsAppDocument(whatsAppDocumentDTO *entity.WhatsAppDocumentDTO) *entity.WhatsAppDocument {
	return &entity.WhatsAppDocument{
		Msisdn:   whatsAppDocumentDTO.Msisdn,
		Message:  whatsAppDocumentDTO.Message,
		Document: whatsAppDocumentDTO.Document,
		FileName: whatsAppDocumentDTO.FileName,
		FileType: whatsAppDocumentDTO.FileType,
	}
}
