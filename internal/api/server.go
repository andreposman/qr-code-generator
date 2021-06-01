package api

import (
	"net/http"

	"github.com/andreposman/qr-code-generator/internal/qrcode/controller"
	"github.com/andreposman/qr-code-generator/internal/qrcode/model"
	"github.com/andreposman/qr-code-generator/internal/qrcode/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		Router: router,
	}
}

func InitAPI(s *Server) {
	router := s.Router

	// Default health-check
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/qrcode/create", func(c *gin.Context) {
		var JSONReq model.QRCodeParam
		err := c.BindJSON(&JSONReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err2 := controller.CreateQRCode(JSONReq)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		} else {
			stringImage, _ := service.EncodeQRCodeString(JSONReq.FileName)
			c.JSON(http.StatusOK, gin.H{
				"image":          stringImage,
				"online-decoder": "https://www.base64decode.net/base64-image-decoder"})

		}

	})

	router.Run(":9000")
}
