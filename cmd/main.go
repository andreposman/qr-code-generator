package main

import (
	"fmt"
	"qr-code-generator/service/qrcode.go"
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
