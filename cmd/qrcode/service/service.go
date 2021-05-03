package service

import (
	"errors"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateQRCode(data string) (barcode.Barcode, error) {
	if len(data) <= 0 || len(data) > 1100 {
		return nil, fmt.Errorf("the string size should be bigger than zero and smaller than 1101")
	}
	qrCode, err := qr.Encode(data, qr.M, qr.Auto)
	if err != nil {
		return nil, fmt.Errorf("error encoding the QRCode. Reason: %s", err)

	}

	qrCode, err = barcode.Scale(qrCode, 350, 350)
	if err != nil {
		return nil, fmt.Errorf("error scaling the QRCode. Reason: %s", err)
	}

	return qrCode, nil
}

func SaveQRCode(qrCode barcode.Barcode) (bool, error) {
	fileName := "qrCode.png"
	file, err := os.Create(fileName)
	if err != nil {
		return false, errors.New("error while saving the QRCode")
	}
	defer file.Close()

	QRCode, err := encodeToPNG(file, qrCode)
	if err != nil {
		return false, err
	}

	return QRCode, nil
}

func encodeToPNG(file *os.File, qrCode barcode.Barcode) (bool, error) {
	err := png.Encode(file, qrCode)
	if err != nil {
		return false, errors.New("error while encoding the QRCode to PNG")
	} else {
		return true, nil
	}
}
