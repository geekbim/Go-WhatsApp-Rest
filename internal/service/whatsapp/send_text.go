package whatsapp

import (
	"context"
	"errors"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (w *whatsAppService) WhatsAppSendText(ctx context.Context, jid string, rjid types.JID, message string) (string, error) {
	var err error

	if WhatsAppClient[jid] == nil {
		return "", errors.New("WhatsApp Client is not Valid")
	}

	// Make Sure WhatsApp Client is OK
	err = w.WhatsAppIsClientOK(jid)
	if err != nil {
		return "", err
	}

	// Set Chat Presence
	w.WhatsAppComposeStatus(jid, rjid, true, false)
	defer w.WhatsAppComposeStatus(jid, rjid, false, false)

	// Compose WhatsApp Proto
	msgContent := &waE2E.Message{
		Conversation: proto.String(message),
	}

	// Send WhatsApp Message Proto
	resp, err := WhatsAppClient[jid].SendMessage(ctx, rjid, msgContent)
	if err != nil {
		return "", err
	}

	// Return Error WhatsApp Client is not Valid
	return resp.ID, nil
}
