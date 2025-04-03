package entity_test

import (
	"go_wa_rest/domain/entity"
	"go_wa_rest/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppStatusDomain(t *testing.T) {
	whatsAppStatusDTO := testdata.NewWhatsAppStatusDTO()

	type args struct {
		whatsAppStatusDTO *entity.WhatsAppStatusDTO
	}

	type wantResponse struct {
		whatsAppStatus *entity.WhatsAppStatus
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
			name: "NewWhatsAppStatusRead",
			args: args{
				whatsAppStatusDTO: whatsAppStatusDTO,
			},
			wantResponse: wantResponse{
				whatsAppStatus: &entity.WhatsAppStatus{
					MessageId: whatsAppStatusDTO.MessageId,
					Status:    "read",
				},
			},
		},
		{
			name: "NewWhatsAppStatusDelivered",
			args: args{
				whatsAppStatusDTO: whatsAppStatusDTO,
			},
			wantResponse: wantResponse{
				whatsAppStatus: &entity.WhatsAppStatus{
					MessageId: whatsAppStatusDTO.MessageId,
					Status:    "delivered",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := entity.NewWhatsAppStatus(tt.args.whatsAppStatusDTO)
			if err != nil {
				assert.Equal(t, tt.wantErr.err, err.Errors)
				assert.Nil(t, res)
			}
			assert.Equal(t, tt.wantResponse.whatsAppStatus, res)
		})
	}
}
