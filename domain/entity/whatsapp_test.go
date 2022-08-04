package entity_test

import (
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppDomain(t *testing.T) {
	whatsAppDTO := testdata.NewWhatsAppDTO()

	t.Run("NewWhatsApp", func(t *testing.T) {
		res, err := entity.NewWhatsApp(whatsAppDTO)
		assert.Nil(t, err)
		assert.Equal(t, whatsAppDTO.Msisdn, res.Msisdn)
		assert.Equal(t, whatsAppDTO.Message, res.Message)
	})

	t.Run("NewWhatsAppErrMsisdn", func(t *testing.T) {
		whatsAppDTO.Msisdn = ""
		err := errors.New("invalid msisdn")
		expectedErr := []error{
			err,
		}
		res, errEntity := entity.NewWhatsApp(whatsAppDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})

	t.Run("NewWhatsAppErrMessage", func(t *testing.T) {
		whatsAppDTO.Msisdn = "08123456789"
		whatsAppDTO.Message = ""
		err := errors.New("message cannot be empty")
		expectedErr := []error{
			err,
		}
		res, errEntity := entity.NewWhatsApp(whatsAppDTO)

		assert.Equal(t, expectedErr, errEntity.Errors)
		assert.Nil(t, res)
	})
}
