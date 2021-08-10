package controllers

import (
	"advBridge/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	logv "github.com/sirupsen/logrus"
)

type MQTTResponse struct {
	Id          string                 `json:"id" binding:"required"`
	Cmd         string                 `json:"cmd" binding:"required"`
	Body        map[string]interface{} `json:"body" binding:"required"`
	Status      int                    `json:"status" binding:"required"`
	DisplayName string                 `json:"DisplayName" `
	Message     string                 `json:"message, omitempty"`
}

type TopicController struct {
	Messages   chan MQTTResponse
	mqttClient mqtt.Client
}

var rfidMQTTModel = new(models.RFIDMQTTModel)


func (topic *TopicController) Init() {
	logv.Info(" === MQTT init connect === ")
	rfidMQTTModel.ConnectionToRFIDServer()
}

func (topic *TopicController) ReConnectToServer() {
	logv.Info(" === MQTT start reconnect === ")
	rfidMQTTModel.DisconnectionToRFIDServer()
	rfidMQTTModel.ConnectionToRFIDServer()
}

func (topic *TopicController) SendDataToServer() {
	logv.Info(" === MQTT SendDataToServer === ")
	rfidMQTTModel.PublishToRFIDServerTest()
}

// ============= API ===============

