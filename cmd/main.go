package main

import (
	"fmt"

	"github.com/andreposman/qr-code-generator/internal/helpers/mock"
	"github.com/andreposman/qr-code-generator/internal/qrcode/controller"
)

func main() {
	greeting()

	QRCodeData := mock.MockParamForQRCodeGen()
	controller.CreateQRCode(QRCodeData)
}

func greeting() {
	fmt.Println("\n QR CODE GENERATOR\n-------------------")
}
