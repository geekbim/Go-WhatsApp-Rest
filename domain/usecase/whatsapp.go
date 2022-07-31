package usecase

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
)

type WhatsAppUseCase interface {
	GetQr(ctx context.Context) (string, int, *exceptions.CustomerError)
	SendMessage(ctx context.Context, whatsApp *entity.WhatsApp) (*entity.WhatsApp, *exceptions.CustomerError)
	Logout(ctx context.Context) *exceptions.CustomerError
}
