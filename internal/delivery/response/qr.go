package response

type QR struct {
	QrImage   string `json:"qrImage"`
	QrTimeOut int    `json:"qrTimeOut"`
}

func MapQRToResponse(qrImage string, qrTimeOut int) *QR {
	return &QR{
		QrImage:   qrImage,
		QrTimeOut: qrTimeOut,
	}
}
