package usecase

import (
	"context"
	"go-wa-rest/domain/entity"
	"go-wa-rest/pkg/exceptions"
)

type WhatsAppUseCase interface {
	GetQr(ctx context.Context) (string, int, *exceptions.CustomerError)
	SendMessage(ctx context.Context, whatsApp *entity.WhatsApp) (*entity.WhatsApp, *exceptions.CustomerError)
}
