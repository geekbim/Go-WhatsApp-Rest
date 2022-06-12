package request

type WhatsApp struct {
	Msisdn  string `json:"msisdn"`
	Message string `json:"message"`
}
