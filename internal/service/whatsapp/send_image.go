package whatsapp

import (
	"bytes"
	"context"
	"errors"
	"go_wa_rest/domain/entity"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (w *whatsAppService) WhatsAppSendImage(ctx context.Context, jid string, rjid types.JID, whatsAppImage *entity.WhatsAppImage) (string, error) {
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

	imgThumbEncode := new(bytes.Buffer)

	imageUploaded, err := WhatsAppClient[jid].Upload(ctx, whatsAppImage.Image, whatsmeow.MediaImage)
	if err != nil {
		return "", err
	}

	imageThumbUploaded, err := WhatsAppClient[jid].Upload(ctx, imgThumbEncode.Bytes(), whatsmeow.MediaLinkThumbnail)
	if err != nil {
		return "", err
	}

	// Compose WhatsApp Proto
	msgContent := &waE2E.Message{
		ImageMessage: &waE2E.ImageMessage{
			URL:                 proto.String(imageUploaded.URL),
			DirectPath:          proto.String(imageUploaded.DirectPath),
			Mimetype:            proto.String(whatsAppImage.FileType),
			Caption:             proto.String(whatsAppImage.Message),
			FileLength:          proto.Uint64(imageUploaded.FileLength),
			FileSHA256:          imageUploaded.FileSHA256,
			FileEncSHA256:       imageUploaded.FileEncSHA256,
			MediaKey:            imageUploaded.MediaKey,
			JPEGThumbnail:       imgThumbEncode.Bytes(),
			ThumbnailDirectPath: &imageThumbUploaded.DirectPath,
			ThumbnailSHA256:     imageThumbUploaded.FileSHA256,
			ThumbnailEncSHA256:  imageThumbUploaded.FileEncSHA256,
			ViewOnce:            proto.Bool(false),
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
