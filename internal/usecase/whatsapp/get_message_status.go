package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
	"sync"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type messageStatus struct {
	sync.RWMutex
	StatusMap map[string]string
}

var messageStatuses = &messageStatus{
	StatusMap: make(map[string]string),
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Receipt:
		for _, receipt := range v.MessageIDs {
			messageID := receipt
			status := ""

			switch v.Type {
			case types.ReceiptTypeDelivered:
				status = "delivered"
			case types.ReceiptTypeRead:
				status = "read"
			case types.ReceiptTypePlayed:
				status = "opened"
			default:
				status = "sent"
			}

			messageStatuses.Lock()
			messageStatuses.StatusMap[messageID] = status
			messageStatuses.Unlock()
		}
	}
}

func (interactor *whatsAppInteractor) GetMessageStatus(ctx context.Context, whatsAppStatus *entity.WhatsAppStatus) (*entity.WhatsAppStatus, *exceptions.CustomerError) {
	var multierr *multierror.Error

	if interactor.waClient == nil {
		multierr = multierror.Append(multierr, errors.New("session not found"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	interactor.waClient.AddEventHandler(eventHandler)

	messageStatuses.RLock()
	status, _ := messageStatuses.StatusMap[whatsAppStatus.MessageId]
	messageStatuses.RUnlock()
	if status == "" {
		status = "sent"
	}

	whatsAppStatus.Status = status

	return whatsAppStatus, nil
}
