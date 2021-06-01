package main

import (
	"fmt"

	"github.com/andreposman/qr-code-generator/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	greeting()
	router := gin.Default()
	r := api.NewServer(router)

	// service.ConvertQRCodeString()

	api.InitAPI(r)

	// QRCodeData := mock.MockParamForQRCodeGen()
	// controller.CreateQRCode(QRCodeData)
}

func greeting() {
	fmt.Println("\n QR CODE GENERATOR\n-------------------")
}
