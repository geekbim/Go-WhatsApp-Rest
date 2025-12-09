package request

type WhatsAppContact struct {
	Msisdn    string `json:"msisdn"`
	FullName  string `json:"fullName"`
	FirstName string `json:"firstName"`
}

