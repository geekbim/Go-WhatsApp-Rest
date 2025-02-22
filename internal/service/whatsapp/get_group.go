package whatsapp

import (
	"errors"

	"go.mau.fi/whatsmeow/types"
)

func (w *whatsAppService) WhatsAppGroup(jid string) ([]*types.GroupInfo, error) {
	if WhatsAppClient[jid] != nil {
		var err error

		// Make Sure WhatsApp Client is OK
		err = w.WhatsAppIsClientOK(jid)
		if err != nil {
			return nil, err
		}

		// Send WhatsApp Message Proto
		groups, err := WhatsAppClient[jid].GetJoinedGroups()
		if err != nil {
			return nil, err
		}

		return groups, nil
	}

	// Return Error WhatsApp Client is not Valid
	return nil, errors.New("WhatsApp Client is not Valid")
}
