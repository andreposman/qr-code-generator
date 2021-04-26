package main

import (
	"fmt"

	"github.com/andreposman/qr-code-generator"
)

func main() {
	greeting()
	fileName := "qrcode.png"

	QRCode := qrcode.Generate("123-456")
	qr.SaveQRCode(fileName, QRCode)
}

func greeting() {
	fmt.Println("\n QR CODE GENERATOR\n-------------------\n")
}
