package whatsapp

import (
	"context"
	"errors"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"google.golang.org/protobuf/proto"
)

func (w *whatsAppService) WhatsAppSendText(ctx context.Context, jid string, rjid string, message string) (string, error) {
	if WhatsAppClient[jid] != nil {
		var err error

		// Make Sure WhatsApp Client is OK
		err = w.WhatsAppIsClientOK(jid)
		if err != nil {
			return "", err
		}

		// Compose New Remote JID
		remoteJID := w.WhatsAppComposeJID(rjid)

		// Set Chat Presence
		w.WhatsAppComposeStatus(jid, remoteJID, true, false)
		defer w.WhatsAppComposeStatus(jid, remoteJID, false, false)

		// Compose WhatsApp Proto
		msgId := WhatsAppClient[jid].GenerateMessageID()
		msgContent := &waE2E.Message{
			Conversation: proto.String(message),
		}

		// Send WhatsApp Message Proto
		_, err = WhatsAppClient[jid].SendMessage(ctx, remoteJID, msgContent)
		if err != nil {
			return "", err
		}

		return msgId, nil
	}

	// Return Error WhatsApp Client is not Valid
	return "", errors.New("WhatsApp Client is not Valid")
}
