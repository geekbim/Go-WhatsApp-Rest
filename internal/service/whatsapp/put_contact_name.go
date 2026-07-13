package whatsapp

import (
	"context"
	"errors"
)

func (w *whatsAppService) WhatsAppPutContactName(ctx context.Context, jid string, msisdn string, fullName string, firstName string) error {
	if WhatsAppClient[jid] == nil {
		return errors.New("WhatsApp Client is not Valid")
	}

	if err := w.WhatsAppIsClientOK(jid); err != nil {
		return err
	}

	contactJID := w.WhatsAppComposeJID(msisdn)

	err := WhatsAppClient[jid].Store.Contacts.PutContactName(ctx, contactJID, fullName, firstName)
	if err != nil {
		return err
	}

	return nil
}
