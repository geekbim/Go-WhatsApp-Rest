package service

import (
	"context"

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
	WhatsAppSendText(ctx context.Context, jid string, rjid string, message string) (string, error)
}
