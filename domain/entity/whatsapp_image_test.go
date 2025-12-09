package entity_test

import (
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/testdata"
	"go_wa_rest/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppImageDomain(t *testing.T) {
	whatsAppImageDTO := testdata.NewWhatsAppImageDTO()

	type args struct {
		whatsAppImageDTO *entity.WhatsAppImageDTO
	}

	type wantResponse struct {
		whatsAppImage *entity.WhatsAppImage
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
			name: "NewWhatsAppImage",
			args: args{
				whatsAppImageDTO: whatsAppImageDTO,
			},
			wantResponse: wantResponse{
				whatsAppImage: &entity.WhatsAppImage{
					ChatType: valueobject.NewChatType(whatsAppImageDTO.ChatType),
					Msisdn:   whatsAppImageDTO.Msisdn,
					Image:    whatsAppImageDTO.Image,
					FileName: whatsAppImageDTO.FileName,
					FileType: whatsAppImageDTO.FileType,
				},
			},
		},
		{
			name: "NewWhatsAppImageErrChatType",
			args: args{
				whatsAppImageDTO: &entity.WhatsAppImageDTO{
					ChatType: "",
					Msisdn:   whatsAppImageDTO.Msisdn,
					Image:    whatsAppImageDTO.Image,
					FileName: whatsAppImageDTO.FileName,
					FileType: whatsAppImageDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("chat type cannot be empty"),
				},
			},
		},
		{
			name: "NewWhatsAppImageErrMsisdn",
			args: args{
				whatsAppImageDTO: &entity.WhatsAppImageDTO{
					ChatType: whatsAppImageDTO.ChatType,
					Msisdn:   "",
					Image:    whatsAppImageDTO.Image,
					FileName: whatsAppImageDTO.FileName,
					FileType: whatsAppImageDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("invalid msisdn"),
				},
			},
		},
		{
			name: "NewWhatsAppImageErrImage",
			args: args{
				whatsAppImageDTO: &entity.WhatsAppImageDTO{
					ChatType: whatsAppImageDTO.ChatType,
					Msisdn:   whatsAppImageDTO.Msisdn,
					Image:    nil,
					FileName: whatsAppImageDTO.FileName,
					FileType: whatsAppImageDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("image cannot be empty"),
				},
			},
		},
		{
			name: "NewWhatsAppImageErrFileName",
			args: args{
				whatsAppImageDTO: &entity.WhatsAppImageDTO{
					ChatType: whatsAppImageDTO.ChatType,
					Msisdn:   whatsAppImageDTO.Msisdn,
					Image:    whatsAppImageDTO.Image,
					FileName: "",
					FileType: whatsAppImageDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("file name cannot be empty"),
				},
			},
		},
		{
			name: "NewWhatsAppImageErrFileType",
			args: args{
				whatsAppImageDTO: &entity.WhatsAppImageDTO{
					ChatType: whatsAppImageDTO.ChatType,
					Msisdn:   whatsAppImageDTO.Msisdn,
					Image:    whatsAppImageDTO.Image,
					FileName: whatsAppImageDTO.FileName,
					FileType: "",
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("file type cannot be empty"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := entity.NewWhatsAppImage(tt.args.whatsAppImageDTO)
			if err != nil {
				assert.Equal(t, tt.wantErr.err, err.Errors)
				assert.Nil(t, res)
			}
			assert.Equal(t, tt.wantResponse.whatsAppImage, res)
		})
	}
}
