package service_test

import (
	"errors"
	"testing"

	"github.com/andreposman/qr-code-generator/cmd/helpers"
	"github.com/andreposman/qr-code-generator/cmd/qrcode/service"
	"github.com/stretchr/testify/assert"
)

func TestServiceShouldGenerateValidQRCode(t *testing.T) {
	randomData := helpers.GenerateRandomData(16)
	result, err := service.GenerateQRCode(randomData)
	if err != nil {
		t.Error(err)
	}
	assert.Greater(t, len(result.Content()), 0)
}

func TestServiceShouldSaveValidQRCode(t *testing.T) {
	randomData := helpers.GenerateRandomData(16)
	result, err := service.GenerateQRCode(randomData)
	if err != nil {
		t.Error(err)
	}

	QRCodeCreated, err := service.SaveQRCode(result)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, QRCodeCreated, "QRCode created succesfully")
}

func TestServiceShouldReturnErrorGeneratingInvalidQRCode(t *testing.T) {
	randomData := helpers.GenerateRandomData(0)
	_, err := service.GenerateQRCode(randomData)
	if err != nil {
		err = errors.New("forced error")
	}

	assert.Error(t, err)
}
