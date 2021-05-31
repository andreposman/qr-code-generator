package controller_test

import (
	"testing"
	"time"

	"github.com/andreposman/qr-code-generator/internal/helpers"
	"github.com/andreposman/qr-code-generator/internal/qrcode/controller"
	"github.com/andreposman/qr-code-generator/internal/qrcode/model"
	"github.com/stretchr/testify/assert"
)

func TestControllerShouldGenerateAndSaveValidQRCode(t *testing.T) {
	mockData := model.QRCodeParam{
		Data:     helpers.GenerateRandomData(128),
		FileName: "controller_unit_testing" + time.Now().Format("01-02-2006 15:04:05"),
	}

	err := controller.CreateQRCode(mockData)

	assert.Nil(t, err)
}
