package usecase

import (
	"context"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"

	"go.mau.fi/whatsmeow/types"
)

type WhatsAppUseCase interface {
	GetQr(ctx context.Context) (string, int, *exceptions.CustomerError)
	GetQrV2(ctx context.Context, jid string) (string, int, *exceptions.CustomerError)
	SendMessage(ctx context.Context, whatsApp *entity.WhatsApp) (*entity.WhatsApp, *exceptions.CustomerError)
	SendMessageV2(ctx context.Context, whatsApp *entity.WhatsApp, jid string) (*entity.WhatsApp, *exceptions.CustomerError)
	SendDocument(ctx context.Context, whatsAppDocument *entity.WhatsAppDocument) (*entity.WhatsAppDocument, *exceptions.CustomerError)
	SendDocumentV2(ctx context.Context, whatsAppDocument *entity.WhatsAppDocument, jid string) (*entity.WhatsAppDocument, *exceptions.CustomerError)
	GetGroup(ctx context.Context) ([]*types.GroupInfo, *exceptions.CustomerError)
	GetGroupV2(ctx context.Context, jid string) ([]*types.GroupInfo, *exceptions.CustomerError)
	Logout(ctx context.Context) *exceptions.CustomerError
	LogoutV2(ctx context.Context, jid string) *exceptions.CustomerError
}
