package apiForms

/**
 * @apiDefine MQTTServerTestDataValidate
 * @apiParam {String} mqttConnectionString mqttConnectionString <a style="color:red">[required]</a>.
 * @apiParam {String} RFIDServerUsername RFIDServerUsername <a style="color:red">[required]</a>.
 * @apiParam {String} RFIDServerPassword RFIDServerPassword <a style="color:red">[required]</a>.
* @apiParamExample {json} Request-Example:
{
	"mqttConnectionString": "tcp://104.215.147.159:1883",
	"RFIDServerPassword": "1JFoR3YbyGaGfNGPGg19Flqzy",
	"RFIDServerUsername": "ec1aceb8-88aa-4b60-8cff-4e8e1cae9e5f:e325b491-edc1-4019-a4e8-675b7c80852c",
}
*/
type MQTTServerTestDataValidate struct {
	MQTTConnectionString string `json:"mqttConnectionString" binding:"required"`
	RFIDServerUsername   string `json:"RFIDServerUsername" binding:"required"`
	RFIDServerPassword   string `json:"RFIDServerPassword" binding:"required"`
}
