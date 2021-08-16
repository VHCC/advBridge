package controllers

import (
	"advBridge/apiForms"
	"advBridge/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
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

/**
@api {POST} /api/v1//mqttServer/connectTest MQTT Server Connection Test
@apiDescription MQTT Server Connection Test
@apiversion 0.0.1
@apiGroup 008 MQTT Server
@apiName MQTT Server Connection Test

@apiUse MQTTServerTestDataValidate

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 2001:CONNECT_ERROR </br>
* @apiSuccess     {String}  message  錯誤訊息
*
* @apiUse HRServerResponse_Success
* @apiUse UserResponse_Invalid_parameter
* @apiUse HRServerResponse_Connect_Err
* @apiUse UserResponse_user_token_invalid
*/
func (topic *TopicController) ConnectTest(c *gin.Context) {
	var data apiForms.MQTTServerTestDataValidate

	// formData validation
	if c.ShouldBind(&data) != nil {
		logv.Error("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	err := rfidMQTTModel.TestConnectionToRFIDServer(data.MQTTConnectionString, data.RFIDServerUsername, data.RFIDServerPassword)
	if err != nil {
		if err != nil {
			c.JSON(200, gin.H{"code": 2001, "message": "CONNECT_ERROR, " + err.Error()})
			logModel.WriteLog(models.EVENT_TYPE_RFID_SERVER_CONNECT_FAIL, data.RFIDServerUsername, err.Error(), nil)
			c.Abort()
			return
		}
	}

	logModel.WriteLog(models.EVENT_TYPE_RFID_SERVER_CONNECT_SUCCESS, data.RFIDServerUsername, "SUCCESS", nil)

	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS"})
}
