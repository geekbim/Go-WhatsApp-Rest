package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
)

func (w *whatsAppService) WhatsAppMessageStatus(ctx context.Context, jid string, messageId string) (*entity.WhatsAppStatus, error) {
	if WhatsAppClient[jid] != nil {
		var err error

		// Make Sure WhatsApp Client is OK
		err = w.WhatsAppIsClientOK(jid)
		if err != nil {
			return nil, err
		}

		WhatsAppClient[jid].AddEventHandler(eventHandler)

		MessageStatuses.RLock()
		status, _ := MessageStatuses.StatusMap[messageId]
		MessageStatuses.RUnlock()
		if status == "" {
			status = "sent"
		}

		whatsAppStatus := entity.WhatsAppStatus{
			MessageId: messageId,
			Status:    status,
		}

		return &whatsAppStatus, nil
	}

	// Return Error WhatsApp Client is not Valid
	return nil, errors.New("WhatsApp Client is not Valid")
}
