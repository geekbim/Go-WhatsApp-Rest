package request

type WhatsApp struct {
	ChatType string `json:"chatType"`
	Msisdn   string `json:"msisdn"`
	Message  string `json:"message"`
}
