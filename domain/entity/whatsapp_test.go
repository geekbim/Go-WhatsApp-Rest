package entity_test

import (
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/testdata"
	"go_wa_rest/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppDomain(t *testing.T) {
	whatsAppDTO := testdata.NewWhatsAppDTO()

	type args struct {
		whatsAppDTO *entity.WhatsAppDTO
	}

	type wantResponse struct {
		whatsApp *entity.WhatsApp
	}

	type wantErr struct {
		err []error
	}

	tests := []struct {
		name         string
		args         args
		wantResponse wantResponse
		wantErr      wantErr
	}{
		{
			name: "NewWhatsApp",
			args: args{
				whatsAppDTO: whatsAppDTO,
			},
			wantResponse: wantResponse{
				whatsApp: &entity.WhatsApp{
					ChatType: valueobject.NewChatType(whatsAppDTO.ChatType),
					Msisdn:   whatsAppDTO.Msisdn,
					Message:  whatsAppDTO.Message,
				},
			},
		},
		{
			name: "NewWhatsAppErrChatType",
			args: args{
				whatsAppDTO: &entity.WhatsAppDTO{
					ChatType: "",
					Msisdn:   whatsAppDTO.Msisdn,
					Message:  whatsAppDTO.Message,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("chat type cannot be empty"),
				},
			},
		},
		{
			name: "NewWhatsAppErrMsisdn",
			args: args{
				whatsAppDTO: &entity.WhatsAppDTO{
					ChatType: whatsAppDTO.ChatType,
					Msisdn:   "",
					Message:  whatsAppDTO.Message,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("invalid msisdn"),
				},
			},
		},
		{
			name: "NewWhatsAppErrMessage",
			args: args{
				whatsAppDTO: &entity.WhatsAppDTO{
					ChatType: whatsAppDTO.ChatType,
					Msisdn:   whatsAppDTO.Msisdn,
					Message:  "",
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("message cannot be empty"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := entity.NewWhatsApp(tt.args.whatsAppDTO)
			if err != nil {
				assert.Equal(t, tt.wantErr.err, err.Errors)
				assert.Nil(t, res)
			}
			assert.Equal(t, tt.wantResponse.whatsApp, res)
		})
	}
}
