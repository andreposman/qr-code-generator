package controller

import (
	"fmt"

	"github.com/andreposman/qr-code-generator/cmd/qrcode/service"
	"github.com/boombuler/barcode"
)

func Generate(data string) barcode.Barcode {
	generatedQRCode, _ := service.GenerateQRCode(data)

	return generatedQRCode
}

func Save(generatedQRCode barcode.Barcode) {

	qrCode, _ := service.SaveQRCode(generatedQRCode)

	if qrCode {
		fmt.Println("\nThe QRCode was successfully generated")
	}
}
