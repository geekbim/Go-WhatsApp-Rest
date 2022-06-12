package response_test

import (
	"go-wa-rest/internal/delivery/response"
	"go-wa-rest/testdata"
	"testing"

	"github.com/go-playground/assert"
)

func TestWhatsAppResponse(t *testing.T) {
	whatsAppDTO := testdata.NewWhatsAppDTO()
	whatsApp := testdata.NewWhatsApp(whatsAppDTO)

	res := response.MapWhatsAppDomainToResponse(whatsApp)

	assert.Equal(t, whatsApp.Msisdn, res.Msisdn)
	assert.Equal(t, whatsApp.Message, res.Message)
}
