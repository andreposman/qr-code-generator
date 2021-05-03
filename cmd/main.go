package main

import (
	"fmt"

	"github.com/andreposman/qr-code-generator/cmd/helpers"
	"github.com/andreposman/qr-code-generator/cmd/qrcode/controller"
)

func main() {
	greeting()

	data := helpers.GenerateRandomData(16)

	fmt.Println(data)

	qrCode := controller.Generate(data)
	controller.Save(qrCode)
}

func greeting() {
	fmt.Println("\n QR CODE GENERATOR\n-------------------")
}
