package whatsapp

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/valueobject"

	"github.com/hashicorp/go-multierror"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (interactor *whatsAppInteractor) SendDocument(ctx context.Context, whatsAppDocument *entity.WhatsAppDocument) (*entity.WhatsAppDocument, *exceptions.CustomerError) {
	var (
		multierr  *multierror.Error
		remoteJID types.JID
	)

	switch whatsAppDocument.ChatType.GetValue() {
	case valueobject.Private:
		remoteJID = interactor.whatsAppService.WhatsAppComposeJID(whatsAppDocument.Msisdn)
	case valueobject.Group:
		remoteJID = types.NewJID(whatsAppDocument.Msisdn, types.GroupServer)
	}

	if interactor.waClient == nil {
		multierr = multierror.Append(multierr, errors.New("session not found"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	fileUploaded, err := interactor.waClient.Upload(ctx, whatsAppDocument.Document, whatsmeow.MediaDocument)
	if err != nil {
		multierr = multierror.Append(multierr, errors.New("error while uploading media to whatsapp server"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	msgExtra := whatsmeow.SendRequestExtra{
		ID: interactor.waClient.GenerateMessageID(),
	}
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

	_, err = interactor.waClient.SendMessage(ctx, remoteJID, msgContent, msgExtra)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return whatsAppDocument, nil
}
