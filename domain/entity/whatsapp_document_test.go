package entity_test

import (
	"errors"
	"go_wa_rest/domain/entity"
	"go_wa_rest/testdata"
	"go_wa_rest/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppDocumentDomain(t *testing.T) {
	whatsAppDocumentDTO := testdata.NewWhatsAppDocumentDTO()

	type args struct {
		whatsAppDocumentDTO *entity.WhatsAppDocumentDTO
	}

	type wantResponse struct {
		whatsAppDocument *entity.WhatsAppDocument
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
			name: "NewWhatsAppDocument",
			args: args{
				whatsAppDocumentDTO: whatsAppDocumentDTO,
			},
			wantResponse: wantResponse{
				whatsAppDocument: &entity.WhatsAppDocument{
					ChatType: valueobject.NewChatType(whatsAppDocumentDTO.ChatType),
					Msisdn:   whatsAppDocumentDTO.Msisdn,
					Document: whatsAppDocumentDTO.Document,
					FileName: whatsAppDocumentDTO.FileName,
					FileType: whatsAppDocumentDTO.FileType,
				},
			},
		},
		{
			name: "NewWhatsAppDocumentErrChatType",
			args: args{
				whatsAppDocumentDTO: &entity.WhatsAppDocumentDTO{
					ChatType: "",
					Msisdn:   whatsAppDocumentDTO.Msisdn,
					Document: whatsAppDocumentDTO.Document,
					FileName: whatsAppDocumentDTO.FileName,
					FileType: whatsAppDocumentDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("chat type cannot be empty"),
				},
			},
		},
		{
			name: "NewWhatsAppDocumentErrMsisdn",
			args: args{
				whatsAppDocumentDTO: &entity.WhatsAppDocumentDTO{
					ChatType: whatsAppDocumentDTO.ChatType,
					Msisdn:   "",
					Document: whatsAppDocumentDTO.Document,
					FileName: whatsAppDocumentDTO.FileName,
					FileType: whatsAppDocumentDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("invalid msisdn"),
				},
			},
		},
		{
			name: "NewWhatsAppDocumentErrDocument",
			args: args{
				whatsAppDocumentDTO: &entity.WhatsAppDocumentDTO{
					ChatType: whatsAppDocumentDTO.ChatType,
					Msisdn:   whatsAppDocumentDTO.Msisdn,
					Document: nil,
					FileName: whatsAppDocumentDTO.FileName,
					FileType: whatsAppDocumentDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("document cannot be empty"),
				},
			},
		},
		{
			name: "NewWhatsAppDocumentErrFileName",
			args: args{
				whatsAppDocumentDTO: &entity.WhatsAppDocumentDTO{
					ChatType: whatsAppDocumentDTO.ChatType,
					Msisdn:   whatsAppDocumentDTO.Msisdn,
					Document: whatsAppDocumentDTO.Document,
					FileName: "",
					FileType: whatsAppDocumentDTO.FileType,
				},
			},
			wantErr: wantErr{
				err: []error{
					errors.New("file name cannot be empty"),
				},
			},
		},
		{
			name: "NewWhatsAppDocumentErrFileType",
			args: args{
				whatsAppDocumentDTO: &entity.WhatsAppDocumentDTO{
					ChatType: whatsAppDocumentDTO.ChatType,
					Msisdn:   whatsAppDocumentDTO.Msisdn,
					Document: whatsAppDocumentDTO.Document,
					FileName: whatsAppDocumentDTO.FileName,
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
			res, err := entity.NewWhatsAppDocument(tt.args.whatsAppDocumentDTO)
			if err != nil {
				assert.Equal(t, tt.wantErr.err, err.Errors)
				assert.Nil(t, res)
			}
			assert.Equal(t, tt.wantResponse.whatsAppDocument, res)
		})
	}
}
