package whatsapp

import (
	"encoding/base64"
	"time"

	qrCode "github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
)

func (w *whatsAppService) WhatsAppGenerateQR(qrChan <-chan whatsmeow.QRChannelItem) (string, int) {
	if qrChan == nil {
		return "", 0
	}
	timer := time.NewTimer(15 * time.Second)
	defer timer.Stop()
	select {
	case evt, ok := <-qrChan:
		if !ok || evt.Event != "code" || evt.Code == "" {
			return "", 0
		}
		qrPNG, err := qrCode.Encode(evt.Code, qrCode.Medium, 256)
		if err != nil {
			return "", 0
		}
		return base64.StdEncoding.EncodeToString(qrPNG), int(evt.Timeout.Seconds())
	case <-timer.C:
		return "", 0
	}
}
