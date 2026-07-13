package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"log"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (w *whatsAppService) WhatsAppSendDocument(ctx context.Context, jid string, rjid types.JID, whatsAppDocument *entity.WhatsAppDocument) (string, error) {
	var err error

	if WhatsAppClient[jid] == nil {
		InitWhatsAppV2(nil, jid)
	}
	if WhatsAppClient[jid] == nil {
		log.Printf("[whatsapp.send_document] client missing jid=%s rjid=%s file=%s", jid, rjid.String(), whatsAppDocument.FileName)
		return "", errors.New("WhatsApp Client is not Valid")
	}

	// Make Sure WhatsApp Client is OK
	err = w.WhatsAppIsClientOK(jid)
	if err != nil {
		if reconnectErr := w.WhatsAppReconnect(jid); reconnectErr == nil {
			err = w.WhatsAppIsClientOK(jid)
		}
	}
	if err != nil {
		log.Printf("[whatsapp.send_document] client not ready jid=%s rjid=%s file=%s err=%v", jid, rjid.String(), whatsAppDocument.FileName, err)
		return "", err
	}

	// Set Chat Presence
	w.WhatsAppComposeStatus(jid, rjid, true, false)
	defer w.WhatsAppComposeStatus(jid, rjid, false, false)

	fileUploaded, err := WhatsAppClient[jid].Upload(ctx, whatsAppDocument.Document, whatsmeow.MediaDocument)
	if err != nil {
		log.Printf("[whatsapp.send_document] upload error jid=%s rjid=%s file=%s err=%v", jid, rjid.String(), whatsAppDocument.FileName, err)
		return "", err
	}

	mentionValues := whatsAppDocument.Mentions
	if entity.HasMentionAll(mentionValues) {
		groupInfo, err := WhatsAppClient[jid].GetGroupInfo(ctx, rjid)
		if err != nil {
			log.Printf("[whatsapp.send_document] group info error jid=%s rjid=%s file=%s err=%v", jid, rjid.String(), whatsAppDocument.FileName, err)
			return "", err
		}
		log.Printf("[whatsapp.send_document] mention all expanded jid=%s rjid=%s participants=%d file=%s", jid, rjid.String(), len(groupInfo.Participants), whatsAppDocument.FileName)
		for _, participant := range groupInfo.Participants {
			mentionValues = append(mentionValues, participant.JID.String())
		}
	}

	mentions := entity.NormalizeMentionJIDs(mentionValues)
	if len(mentionValues) > 0 {
		log.Printf("[whatsapp.send_document] mentions normalized jid=%s rjid=%s input=%d normalized=%d file=%s", jid, rjid.String(), len(mentionValues), len(mentions), whatsAppDocument.FileName)
	}
	var contextInfo *waE2E.ContextInfo
	if len(mentions) > 0 {
		contextInfo = &waE2E.ContextInfo{MentionedJID: mentions}
	}

	// Compose WhatsApp Proto
	msgContent := &waE2E.Message{
		DocumentMessage: &waE2E.DocumentMessage{
			URL:           proto.String(fileUploaded.URL),
			DirectPath:    proto.String(fileUploaded.DirectPath),
			Mimetype:      proto.String(whatsAppDocument.FileType),
			Caption:       proto.String(whatsAppDocument.Message),
			Title:         proto.String(whatsAppDocument.FileName),
			FileName:      proto.String(whatsAppDocument.FileName),
			FileLength:    proto.Uint64(fileUploaded.FileLength),
			FileSHA256:    fileUploaded.FileSHA256,
			FileEncSHA256: fileUploaded.FileEncSHA256,
			MediaKey:      fileUploaded.MediaKey,
			ContextInfo:   contextInfo,
		},
	}

	// Send WhatsApp Message Proto
	resp, err := WhatsAppClient[jid].SendMessage(ctx, rjid, msgContent)
	if err != nil {
		log.Printf("[whatsapp.send_document] send message error jid=%s rjid=%s file=%s err=%v", jid, rjid.String(), whatsAppDocument.FileName, err)
		return "", err
	}

	log.Printf("[whatsapp.send_document] sent jid=%s rjid=%s file=%s message_id=%s", jid, rjid.String(), whatsAppDocument.FileName, resp.ID)

	// Return Error WhatsApp Client is not Valid
	return resp.ID, nil
}
