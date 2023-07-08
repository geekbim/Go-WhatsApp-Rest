package whatsapp_test

import (
	"context"
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/internal/usecase/whatsapp"
	"go_wa_rest/mocks"
	"go_wa_rest/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendMessageV2(t *testing.T) {
	whatsAppDto := testdata.NewWhatsAppDTO()
	whatsAppEntity := testdata.NewWhatsApp(whatsAppDto)

	type args struct {
		whatsApp *entity.WhatsApp
		jid      string
	}

	tests := []struct {
		name             string
		args             args
		whatsAppSendText funcCall
	}{
		{
			name: "OK",
			args: args{
				whatsApp: whatsAppEntity,
				jid:      "jid",
			},
			whatsAppSendText: funcCall{
				input: []interface{}{
					context.Background(),
					"jid",
					whatsAppEntity.Msisdn,
					whatsAppEntity.Message,
				},
				output: []interface{}{
					"",
					nil,
				},
			},
		},
		{
			name: "ErrWhatsAppSendText",
			args: args{
				whatsApp: whatsAppEntity,
				jid:      "jid",
			},
			whatsAppSendText: funcCall{
				input: []interface{}{
					context.Background(),
					"jid",
					whatsAppEntity.Msisdn,
					whatsAppEntity.Message,
				},
				output: []interface{}{
					"",
					errors.New("WhatsAppSendText"),
				},
			},
		},
	}

	for _, tt := range tests {
		whatsAppService := new(mocks.WhatsAppService)

		whatsAppService.
			On("WhatsAppSendText", tt.whatsAppSendText.input...).
			Return(tt.whatsAppSendText.output...)

		useCase := whatsapp.NewWhatsAppInteractor(nil, whatsAppService)
		res, err := useCase.SendMessageV2(context.Background(), tt.args.whatsApp, tt.args.jid)
		if err != nil {
			assert.NotNil(t, err)
			return
		}
		assert.Equal(t, whatsAppEntity, res)
	}
}
