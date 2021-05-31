package controller

import (
	"fmt"

	"github.com/andreposman/qr-code-generator/internal/qrcode/model"
	"github.com/andreposman/qr-code-generator/internal/qrcode/service"
	"github.com/boombuler/barcode"
)

func CreateQRCode(QRCodeData model.QRCodeParam) error {
	genQRCode, err := generate(QRCodeData.Data)
	if err != nil {
		return err
	}

	errSave := save(genQRCode, QRCodeData.FileName)
	if errSave != nil {
		return err
	} else {
		return nil
	}
}

func generate(data string) (barcode.Barcode, error) {
	generatedQRCode, err := service.GenerateQRCode(data)
	if err != nil {
		return nil, err
	} else {
		return generatedQRCode, nil
	}
}

func save(generatedQRCode barcode.Barcode, fileName string) error {
	QRCode, err := service.SaveQRCode(generatedQRCode, fileName)
	if err != nil {
		return err
	}

	if QRCode {
		fmt.Println("\nThe QRCode was successfully created and saved!")
	}

	return nil
}
