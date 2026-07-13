package whatsapp

import (
	"context"
	"go_wa_rest/pkg/exceptions"
	"go_wa_rest/valueobject"

	"go.mau.fi/whatsmeow/types"
)

func (interactor *whatsAppInteractor) StartTypingV2(ctx context.Context, chatType valueobject.ChatTypeEnum, msisdn string, jid string, isAudio bool) *exceptions.CustomerError {
	var remoteJID types.JID

	switch chatType {
	case valueobject.Private:
		remoteJID = interactor.whatsAppService.WhatsAppComposeJID(msisdn)
	case valueobject.Group:
		remoteJID = types.NewJID(msisdn, types.GroupServer)
	default:
		return &exceptions.CustomerError{
			Status: exceptions.ERRBUSSINESS,
		}
	}

	interactor.whatsAppService.WhatsAppComposeStatus(jid, remoteJID, true, isAudio)

	return nil
}
