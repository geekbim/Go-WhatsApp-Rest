package service

import (
	"context"
	"go_wa_rest/domain/entity"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

type WhatsAppService interface {
	WhatsAppIsClientOK(jid string) error
	WhatsAppComposeJID(jid string) types.JID
	WhatsAppComposeStatus(jid string, rjid types.JID, isComposing bool, isAudio bool)
	WhatsAppDecomposeJID(jid string) string
	WhatsAppGenerateQR(qrChan <-chan whatsmeow.QRChannelItem) (string, int)
	WhatsAppLogin(jid string) (string, int, error)
	WhatsAppLogout(jid string) error
	WhatsAppReconnect(jid string) error
	WhatsAppGroup(jid string) ([]*types.GroupInfo, error)
	WhatsAppContact(jid string) (map[types.JID]types.ContactInfo, error)
	WhatsAppSendText(ctx context.Context, jid string, rjid types.JID, message string) (string, error)
	WhatsAppSendDocument(ctx context.Context, jid string, rjid types.JID, whatsAppDocument *entity.WhatsAppDocument) (string, error)
	WhatsAppSendImage(ctx context.Context, jid string, rjid types.JID, whatsAppImage *entity.WhatsAppImage) (string, error)
	WhatsAppMessageStatus(ctx context.Context, jid string, messageId string) (*entity.WhatsAppStatus, error)
}
