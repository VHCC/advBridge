package controllers

import (
	"github.com/sacOO7/gowebsocket"
	"advBridge/models"
)

type GlobalConfigController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var globalConfigModel = new(models.GlobalConfigModel)
var logModel = new(models.VmsLogModel)


