package service_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/andreposman/qr-code-generator/internal/helpers"
	"github.com/andreposman/qr-code-generator/internal/qrcode/service"
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

	QRCodeCreated, err := service.SaveQRCode(result, "unit_test_service"+time.Now().Format("01-02-2006 15:04:05"))
	if err != nil {
		t.Error(err)
	}
	assert.True(t, QRCodeCreated, "QRCode created successfully")
}

func TestServiceShouldReturnErrorGeneratingInvalidQRCode(t *testing.T) {
	randomData := helpers.GenerateRandomData(0)
	_, err := service.GenerateQRCode(randomData)
	if err != nil {
		err = errors.New("forced error")
	}

	assert.Error(t, err)
}

func TestShouldGenerateValidQRCode(t *testing.T) {
	randomData := helpers.GenerateRandomData(128)
	qrCode, _ := service.GenerateQRCode(randomData)
	fmt.Println(qrCode)

	// assert.True(t, len(qrCode.Content()) > 0)
}
