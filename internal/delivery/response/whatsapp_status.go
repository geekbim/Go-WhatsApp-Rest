package response

import (
	"go_wa_rest/domain/entity"
)

type WhatsAppStatus struct {
	MessageId string `json:"messageId"`
	Status    string `json:"status"`
}

func MapWhatsAppStatusDomainToResponse(whatsAppStatus *entity.WhatsAppStatus) *WhatsAppStatus {
	return &WhatsAppStatus{
		MessageId: whatsAppStatus.MessageId,
		Status:    whatsAppStatus.Status,
	}
}
