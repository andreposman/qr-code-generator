package qrcode

import "io/ioutil"

func Generate(code string) []byte {
	return nil
}

func Save(filenName string, QRCode []byte) {
	ioutil.WriteFile("qrcode.png", QRCode, 0644)
}
