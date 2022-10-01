package whatsapp

import (
	"context"
	"errors"

	"go.mau.fi/whatsmeow"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func WhatsAppSendText(ctx context.Context, jid string, rjid string, message string) (string, error) {
	if WhatsAppClient[jid] != nil {
		var err error

		// Make Sure WhatsApp Client is OK
		err = WhatsAppIsClientOK(jid)
		if err != nil {
			return "", err
		}

		// Compose New Remote JID
		remoteJID := WhatsAppComposeJID(rjid)

		// Set Chat Presence
		WhatsAppComposeStatus(jid, remoteJID, true, false)
		defer WhatsAppComposeStatus(jid, remoteJID, false, false)

		// Compose WhatsApp Proto
		msgId := whatsmeow.GenerateMessageID()
		msgContent := &waproto.Message{
			Conversation: proto.String(message),
		}

		// Send WhatsApp Message Proto
		_, err = WhatsAppClient[jid].SendMessage(ctx, remoteJID, msgId, msgContent)
		if err != nil {
			return "", err
		}

		return msgId, nil
	}

	// Return Error WhatsApp Client is not Valid
	return "", errors.New("WhatsApp Client is not Valid")
}
