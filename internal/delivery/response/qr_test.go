package response_test

import (
	"go_wa_rest/internal/delivery/response"
	"testing"

	"github.com/go-playground/assert"
)

func TestQrResponse(t *testing.T) {
	image := "qrImageBase64"
	timeOut := 60

	res := response.MapQRToResponse(image, timeOut)

	assert.Equal(t, image, res.QrImage)
	assert.Equal(t, timeOut, res.QrTimeOut)
}
