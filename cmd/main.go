package main

import (
	"fmt"

	qrcode "github.com/andreposman/qr-code-generator"
)

func main() {
	greeting()
	fileName := "qrcode.png"

	QRCode := qrcode.Generate("123-456")
	qrcode.SaveQRCode(fileName, QRCode)
}

func greeting() {
	fmt.Println("\n QR CODE GENERATOR\n-------------------\n")
}
