package whatsapp

import (
	"errors"

	"go.mau.fi/whatsmeow/types"
)

func WhatsAppReconnect(jid string) error {
	if WhatsAppClient[jid] != nil {
		// Make Sure WebSocket Connection is Disconnected
		WhatsAppClient[jid].Disconnect()

		// Make Sure Store ID is not Empty
		// To do Reconnection
		if WhatsAppClient[jid] != nil {
			err := WhatsAppClient[jid].Connect()
			if err != nil {
				return err
			}

			// Set WhatsApp Client Presence to Available
			_ = WhatsAppClient[jid].SendPresence(types.PresenceAvailable)

			return nil
		}

		return errors.New("WhatsApp Client Store ID is Empty, Please Re-Login and Scan QR Code Again")
	}

	return errors.New("WhatsApp Client is not Valid")
}
