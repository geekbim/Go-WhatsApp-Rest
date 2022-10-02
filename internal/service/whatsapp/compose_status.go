package whatsapp

import (
	"go.mau.fi/whatsmeow/types"
)

func (w *whatsAppService) WhatsAppComposeStatus(jid string, rjid types.JID, isComposing bool, isAudio bool) {
	// Set Compose Status
	var typeCompose types.ChatPresence
	if isComposing {
		typeCompose = types.ChatPresenceComposing
	} else {
		typeCompose = types.ChatPresencePaused
	}

	// Set Compose Media Audio (Recording) or Text (Typing)
	var typeComposeMedia types.ChatPresenceMedia
	if isAudio {
		typeComposeMedia = types.ChatPresenceMediaAudio
	} else {
		typeComposeMedia = types.ChatPresenceMediaText
	}

	// Send Chat Compose Status
	_ = WhatsAppClient[jid].SendChatPresence(rjid, typeCompose, typeComposeMedia)
}
