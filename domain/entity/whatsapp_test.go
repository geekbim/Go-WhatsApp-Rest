package entity_test

import (
	"go-wa-rest/domain/entity"
	"go-wa-rest/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppDomain(t *testing.T) {
	whatsAppDTO := testdata.NewWhatsAppDTO()

	res := entity.NewWhatsApp(whatsAppDTO)
	assert.Equal(t, whatsAppDTO.Msisdn, res.Msisdn)
	assert.Equal(t, whatsAppDTO.Message, res.Message)
}
