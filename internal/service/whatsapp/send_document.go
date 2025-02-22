package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (w *whatsAppService) WhatsAppSendDocument(ctx context.Context, jid string, rjid types.JID, whatsAppDocument *entity.WhatsAppDocument) (string, error) {
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

	fileUploaded, err := WhatsAppClient[jid].Upload(ctx, whatsAppDocument.Document, whatsmeow.MediaDocument)
	if err != nil {
		return "", err
	}

	// Compose WhatsApp Proto
	msgContent := &waE2E.Message{
		DocumentMessage: &waE2E.DocumentMessage{
			URL:           proto.String(fileUploaded.URL),
			DirectPath:    proto.String(fileUploaded.DirectPath),
			Mimetype:      proto.String(whatsAppDocument.FileType),
			Title:         proto.String(whatsAppDocument.FileName),
			FileName:      proto.String(whatsAppDocument.FileName),
			FileLength:    proto.Uint64(fileUploaded.FileLength),
			FileSHA256:    fileUploaded.FileSHA256,
			FileEncSHA256: fileUploaded.FileEncSHA256,
			MediaKey:      fileUploaded.MediaKey,
		},
	}

	// Send WhatsApp Message Proto
	resp, err := WhatsAppClient[jid].SendMessage(ctx, rjid, msgContent)
	if err != nil {
		return "", err
	}

	// Return Error WhatsApp Client is not Valid
	return resp.ID, nil
}
