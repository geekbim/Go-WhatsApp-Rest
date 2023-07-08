package whatsapp_test

import (
	"context"
	"errors"
	"go_wa_rest/internal/usecase/whatsapp"
	"go_wa_rest/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type funcCall struct {
	input  []interface{}
	output []interface{}
}

func Test_LogoutV2(t *testing.T) {
	type args struct {
		jid string
	}

	tests := []struct {
		name           string
		args           args
		whatsAppLogout funcCall
	}{
		{
			name: "OK",
			args: args{
				jid: "jid",
			},
			whatsAppLogout: funcCall{
				input: []interface{}{
					"jid",
				},
				output: []interface{}{
					nil,
				},
			},
		},
		{
			name: "ErrWhatsAppLogout",
			args: args{
				jid: "jid",
			},
			whatsAppLogout: funcCall{
				input: []interface{}{
					"jid",
				},
				output: []interface{}{
					errors.New("ErrWhatsAppLogout"),
				},
			},
		},
	}

	for _, tt := range tests {
		whatsAppService := new(mocks.WhatsAppService)

		whatsAppService.
			On("WhatsAppLogout", tt.whatsAppLogout.input...).
			Return(tt.whatsAppLogout.output...)

		useCase := whatsapp.NewWhatsAppInteractor(nil, whatsAppService)
		err := useCase.LogoutV2(context.Background(), tt.args.jid)
		if err != nil {
			assert.NotNil(t, err)
			return
		}
		assert.Nil(t, err)
	}
}
