package helpers

import (
	"github.com/thanhpk/randstr"
)

func GenerateRandomData(size int) string {
	token := randstr.Hex(size) // generate 128-bit hex string

	return token
}
