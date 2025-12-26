package request

type WhatsApp struct {
	ChatType string `json:"chatType"`
	Msisdn   string `json:"msisdn"`
	Message  string `json:"message"`
}

type WhatsAppTyping struct {
	ChatType string `json:"chatType"`
	Msisdn   string `json:"msisdn"`
	IsAudio  bool   `json:"isAudio"`
}
