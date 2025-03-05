package whatsapp

import (
	"bytes"
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/valueobject"

	"github.com/hashicorp/go-multierror"
	"github.com/sunshineplan/imgconv"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (interactor *whatsAppInteractor) SendImage(ctx context.Context, whatsAppImage *entity.WhatsAppImage) (*entity.WhatsAppImage, *exceptions.CustomerError) {
	var (
		multierr  *multierror.Error
		remoteJID types.JID
	)

	switch whatsAppImage.ChatType.GetValue() {
	case valueobject.Private:
		remoteJID = interactor.whatsAppService.WhatsAppComposeJID(whatsAppImage.Msisdn)
	case valueobject.Group:
		remoteJID = types.NewJID(whatsAppImage.Msisdn, types.GroupServer)
	}

	if interactor.waClient == nil {
		multierr = multierror.Append(multierr, errors.New("session not found"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	imgThumbDecode, err := imgconv.Decode(bytes.NewReader(whatsAppImage.Image))
	if err != nil {
		multierr = multierror.Append(multierr, errors.New("Error While Decoding Thumbnail Image Stream"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	imgThumbEncode := new(bytes.Buffer)

	err = imgconv.Write(imgThumbEncode,
		imgconv.Resize(imgThumbDecode, &imgconv.ResizeOption{Width: 72}),
		&imgconv.FormatOption{Format: imgconv.JPEG})
	if err != nil {
		multierr = multierror.Append(multierr, errors.New("Error While Encoding Thumbnail Image Stream"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	imageUploaded, err := interactor.waClient.Upload(ctx, whatsAppImage.Image, whatsmeow.MediaImage)
	if err != nil {
		multierr = multierror.Append(multierr, errors.New("error while uploading media to whatsapp server"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	imageThumbUploaded, err := interactor.waClient.Upload(ctx, imgThumbEncode.Bytes(), whatsmeow.MediaLinkThumbnail)
	if err != nil {
		multierr = multierror.Append(multierr, errors.New("Error while Uploading Image Thumbnail to WhatsApp Server"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

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

	_, err = interactor.waClient.SendMessage(ctx, remoteJID, msgContent)
	if err != nil {
		multierr = multierror.Append(multierr, err)
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multierr,
		}
	}

	return whatsAppImage, nil
}
