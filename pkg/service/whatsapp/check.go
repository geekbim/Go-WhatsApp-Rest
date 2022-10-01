package whatsapp

import "errors"

func WhatsAppIsClientOK(jid string) error {
	// Make Sure WhatsApp Client is Connected
	if !WhatsAppClient[jid].IsConnected() {
		return errors.New("WhatsApp Client is not Connected")
	}

	// Make Sure WhatsApp Client is Logged In
	if !WhatsAppClient[jid].IsLoggedIn() {
		return errors.New("WhatsApp Client is not Logged In")
	}

	return nil
}
