package controllers

import (
	"advBridge/apiForms"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sacOO7/gowebsocket"
	"advBridge/models"
	"gopkg.in/mgo.v2/bson"
)

type GlobalConfigController struct {
	SessionID string
	//Messages chan frs.FRSWSResponse
	Socket gowebsocket.Socket
}

var globalConfigModel = new(models.GlobalConfigModel)
var logModel = new(models.VmsLogModel)


/**
@api {POST} /api/v1/serverConfig/getConfig  取得 Server Configuration 資訊
@apiDescription Server Configuration
@apiversion 0.0.1
@apiGroup 005 SERVER CONFIG
@apiName Get Server Configuration

@apiUse GlobalConfigGetStructure

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonObject}  serverConfig  Server Configuration
*
* @apiUse ServerConfigResponse_Success
* @apiUse CommonResponse_Invalid_parameter
* @apiUse UserResponse_user_token_invalid
*/
func (globalConfigC *GlobalConfigController) GetGlobalConfig(c *gin.Context) {
	var data apiForms.GetGlobalConfigDataValidate

	// formData validation
	if c.ShouldBind(&data) != nil {
		fmt.Println("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	checkResult, queryUser := userModel.UserTokenCheck(data.UserToken)
	_ = queryUser
	switch checkResult {
	case 1001:
		c.JSON(200, gin.H{"code": 1001, "message": "USER_TOKEN_INVALID"})
		c.Abort()
		return
	}

	serverConfig, err := globalConfigModel.FindFromDB()
	if err != nil {
		fmt.Println("FindFromDB err:> ", err)
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS", "serverConfig": serverConfig})
}

/**
@api {POST} /api/v1/serverConfig/updateConfig 更新 Server Configuration
@apiDescription Server Configuration
@apiversion 0.0.1
@apiGroup 005 SERVER CONFIG
@apiName Update Server Configuration

@apiUse GlobalConfigUpdateStructure

* @apiSuccess     {Number} code  錯誤代碼 </br>
*                 0:SUCCESS (成功) </br>
*                 1:INVALID_PARAMETERS (參數缺少或錯誤) </br>
*				  1001:USER_TOKEN_INVALID (userToken invalid) </br>
*                 11099:OPERATION_FAIL  </br>
* @apiSuccess     {String}  message  錯誤訊息
* @apiSuccess     {JsonObject}  serverConfig  Server Configuration
*
* @apiUse ServerConfigResponse_Success
* @apiUse CommonResponse_Invalid_parameter
* @apiUse Response_Operation_Fail
* @apiUse UserResponse_user_token_invalid
*/
func (globalConfigC *GlobalConfigController) UpdateGlobalConfig(c *gin.Context) {
	var data apiForms.UpdateGlobalConfigDataValidate

	// formData validation
	if c.ShouldBind(&data) != nil {
		fmt.Println("ShouldBind err:> ", c.Errors)
		c.JSON(200, gin.H{"code": 1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	checkResult, queryUser := userModel.UserTokenCheck(data.UserToken)
	_ = queryUser
	switch checkResult {
	case 1:
	case 2:
	case 1001:
		c.JSON(200, gin.H{"code": 1001, "message": "USER_TOKEN_INVALID"})
		c.Abort()
		return
	}

	serverConfig, err := globalConfigModel.UpdateToDB(data)
	if err != nil {
		fmt.Println("UpdateToDB err:> ", err)
		c.JSON(200, gin.H{"code": 11099, "message": "OPERATION_FAIL, " + err.Error()})
		c.Abort()
		return
	}

	res2B, _ := json.Marshal(data)
	var queryRequest map[string]interface{}
	_ = bson.UnmarshalJSON([]byte(string(res2B)), &queryRequest)
	if data.Bundle["log_retention"] != nil || data.Bundle["checkin_retention"] != nil || data.Bundle["snapshot_retention"] != nil{
		err = logModel.WriteLog(models.EVENT_TYPE_RETENTION_UPDATE, queryUser.AccountID, "SUCCESS", queryRequest)
	}

	if data.Bundle["smtp"] != nil || data.Bundle["port"] != nil || data.Bundle["enablessltls"] != nil{
		err = logModel.WriteLog(models.EVENT_TYPE_SMTP_UPDATE, queryUser.AccountID, "SUCCESS", queryRequest)
	}

	c.JSON(200, gin.H{"code": 0, "message": "SUCCESS", "serverConfig": serverConfig})
}

