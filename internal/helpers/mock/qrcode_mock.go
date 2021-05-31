package mock

import (
	"time"

	"github.com/andreposman/qr-code-generator/internal/qrcode/model"
)

func MockParamForQRCodeGen() model.QRCodeParam {
	data := "www.github.com/andreposman"
	fileName := "qrcode_" + time.Now().Format("01-02-2006 15:04:05")

	QRCodeData := model.QRCodeParam{
		Data:     data,
		FileName: fileName,
	}

	return QRCodeData
}
