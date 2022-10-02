package whatsapp

import (
	"strings"

	"go.mau.fi/whatsmeow/types"
)

func (w *whatsAppService) WhatsAppComposeJID(jid string) types.JID {
	// Decompose JID First Before Recomposing
	jid = w.WhatsAppDecomposeJID(jid)

	// Check if JID Contains '-' Symbol
	if strings.ContainsRune(jid, '-') {
		// Check if the JID is a Group ID
		if len(strings.SplitN(jid, "-", 2)) == 2 {
			// Return JID as Group Server (@g.us)
			return types.NewJID(jid, types.GroupServer)
		}
	}

	// Return JID as Default User Server (@s.whatsapp.net)
	return types.NewJID(jid, types.DefaultUserServer)
}
