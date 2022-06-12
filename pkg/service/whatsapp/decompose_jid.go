package whatsapp

import "strings"

func WhatsAppDecomposeJID(jid string) string {
	// Check if JID Contains '@' Symbol
	if strings.ContainsRune(jid, '@') {
		// Split JID Based on '@' Symbol
		// and Get Only The First Section Before The Symbol
		buffers := strings.Split(jid, "@")
		jid = buffers[0]
	}

	// Check if JID First Character is '+' Symbol
	if jid[0] == '+' {
		// Remove '+' Symbol from JID
		jid = jid[1:]
	}

	return jid
}
