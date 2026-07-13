package whatsapp

import (
	"context"
	"errors"

	"go.mau.fi/whatsmeow/types"
)

func (w *whatsAppService) WhatsAppContact(jid string) (map[types.JID]types.ContactInfo, error) {
	if WhatsAppClient[jid] != nil {
		var err error

		// Make Sure WhatsApp Client is OK
		err = w.WhatsAppIsClientOK(jid)
		if err != nil {
			return nil, err
		}

		// Send WhatsApp Message Proto
		contacts, err := WhatsAppClient[jid].Store.Contacts.GetAllContacts(context.Background())
		if err != nil {
			return nil, err
		}

		return contacts, nil
	}

	// Return Error WhatsApp Client is not Valid
	return nil, errors.New("WhatsApp Client is not Valid")
}
