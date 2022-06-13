package entity_test

import (
	"go_wa_rest/domain/entity"
	"go_wa_rest/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppDomain(t *testing.T) {
	whatsAppDTO := testdata.NewWhatsAppDTO()

	res := entity.NewWhatsApp(whatsAppDTO)
	assert.Equal(t, whatsAppDTO.Msisdn, res.Msisdn)
	assert.Equal(t, whatsAppDTO.Message, res.Message)
}
